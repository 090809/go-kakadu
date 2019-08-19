package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	tb "github.com/090809/telebot"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"go-kakadu/handlers"

)

func init() {
	if err := godotenv.Load(); err != nil {
		if _, err := os.Stat("./.env.example"); os.IsExist(err) {
			log.Fatal("File .env.example exists! Fill it and rename to .env, before we go")
		}
		log.Fatalf("Error .env loading: %v", err.Error())
	} else {
		log.Info(".env file loaded")
	}

	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
}

func main() {
	if seedFlag := flag.Bool("seed", false, "Perform seeding on launch"); *seedFlag {
		seed()
	}

	logger := log.New()

	client := getDBClient(logger)
	bot := getTbBot()

	defHandler := handlers.NewHandler(bot, client, logger)

	bot.Handle("/start", handlers.StartHandle(defHandler))
	bot.Handle(tb.OnText, handlers.TextHandle(defHandler))

	bot.Start()
}

func getDBClient(logger *log.Logger) *sql.DB {
	var (
		clientConfig string
		driver       = os.Getenv("DRIVER")
		driverHost   = os.Getenv("DRIVER_HOST")
		driverUser   = os.Getenv("DRIVER_USER")
		driverPass   = os.Getenv("DRIVER_PASS")
		driverDb     = os.Getenv("DRIVER_DB")
	)

	switch driver {
	case "mysql":
		clientConfig = fmt.Sprintf("%v:%v@tcp(%v)/%v", driverUser, driverPass, driverHost, driverDb)
	default:
		clientConfig = fmt.Sprintf("%v:%v@%v/%v", driverUser, driverPass, driverHost, driverDb)
	}

	log.Infof("Connecting to DB with conf: %v", clientConfig)

	client, e := sql.Open(driver, clientConfig)
	if e != nil {
		log.Errorln(e)
		log.Fatalf("Can't load DB with this data: %v", clientConfig)
	}

	if err := client.Ping(); err != nil {
		log.Fatal(err)
	}

	return client
}

func getTbBot() *tb.Bot {
	var (
		listenURL = os.Getenv("URL")
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   listenURL + ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	settings := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	bot, err := tb.NewBot(settings)
	if err != nil {
		log.Fatalln(err)
	}
	return bot
}
