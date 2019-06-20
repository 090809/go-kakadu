package main

import (
	"log"
	"os"

	//"context"
	//"fmt"
	// "net/http"

	// "github.com/gin-gonic/gin"
	tb "github.com/090809/telebot"
	"github.com/joho/godotenv"
	//"go-kakadu/generated/prisma-client"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if _, err := os.Stat("./.env.example"); os.IsNotExist(err) {
			log.Fatalln("File .env.example exists! Rename it to .env, before we go")
		}
		log.Fatalf("Error .env loading: %v", err.Error())
	}

	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		_, err := b.Send(m.Sender, "Hi!")
		if err != nil {
			log.Fatalln(err)
		}
	})

	b.Start()
}
