package main

import (
	"log"
	"net/http"

	"auto-emails/app"
	c "auto-emails/configuration"
	"auto-emails/helper"

	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	dbMS := app.ConnectDatabaseMS(configuration.UserMS, configuration.HostMS, configuration.PasswordMS, configuration.PortDBMS, configuration.DbMS)
	dbMY := app.ConnectDatabaseMY(configuration.UserMY, configuration.HostMY, configuration.PasswordMY, configuration.PortDBMY, configuration.DbMY)

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(dbMS, dbMY, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
