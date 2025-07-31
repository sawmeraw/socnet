package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sawmeraw/gogo/internal/db"
	"github.com/sawmeraw/gogo/internal/env"
	"github.com/sawmeraw/gogo/internal/mailer"
	"github.com/sawmeraw/gogo/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			Socnet API
//	@version		0.0.1
//	@description	API for socnet, a socialnetwork app
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:    env.GetString("ENV", "development"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:3000"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3, //3 days
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_APIKEY", ""),
			},
			fromEmail: env.GetString("FROM_EMAIL", ""),
		},
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:4000"),
	}

	//Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
	//Database init
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)
	mailer, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail)
	if err != nil {
		log.Fatal(err)
	}
	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))

}
