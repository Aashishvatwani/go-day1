package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Aashishvatwani/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{session: s}
}
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	old := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("mongo-golang").C("users").FindId(old).One(&u); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	old := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(old); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "deleted the content ")

}
