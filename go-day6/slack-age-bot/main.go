package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommadEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommadEvents(bot.CommandEvents())
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB command",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year")
				return
			}
			age := 2025 - yob
			if err != nil {
				response.Reply("Invalid year")
				return
			}
			response.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
