package main

import (
	"fmt"
	"github.com/Aashishvatwani/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {

		fmt.Println("mongo not connected", err)
		return nil
	} else {
		fmt.Println("mongo connected")
	}
	return session
}
