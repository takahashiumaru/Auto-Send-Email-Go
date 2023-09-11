package main

import (
	"log"
	"net/http"

	"auto-emails/app"
	c "auto-emails/configuration"
	"auto-emails/helper"

	"github.com/go-playground/validator/v10"
	// Import Gorilla Mux router
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

	// Create a handler function for your routes
	router := app.NewRouter(dbMS, dbMY, validate)

	// Define your routes using http.HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	})

	// Create a server and listen on the desired port
	port := configuration.Port
	server := &http.Server{
		Addr: ":" + port,
	}

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
