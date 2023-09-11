package main

import (
	"log"
	"net/http"

	"auto-emails/app"
	c "auto-emails/configuration"
	"auto-emails/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux" // Import Gorilla Mux router
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	dbMS := app.ConnectDatabaseMS(configuration.UserMS, configuration.HostMS, configuration.PasswordMS, configuration.PortDBMS, configuration.DbMS)
	dbMY := app.ConnectDatabaseMY(configuration.UserMY, configuration.HostMY, configuration.PasswordMY, configuration.PortDBMY, configuration.DbMY)

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(dbMS, dbMY, validate)

	// Create a Gorilla Mux router and add your router as a subrouter
	r := mux.NewRouter()
	r.Handle("/", router)

	// Create a server and listen on the desired port
	port := configuration.Port
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r, // Use Gorilla Mux router as the main handler
	}

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
