package main

import (
	"log"
	"net/http"

	"auto-emails/app"
	"auto-emails/auth"
	c "auto-emails/configuration"
	"auto-emails/controller"
	"auto-emails/helper"
	"auto-emails/repository"
	"auto-emails/service"

	"github.com/gin-gonic/gin" // Import Gin
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

	// Create a Gin router
	router := NewRouter(dbMS, dbMY, validate)

	// Create a server and listen on the desired port
	port := configuration.Port
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func NewRouter(dbMS *gorm.DB, dbMY *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.Default() // Use Gin in "release" mode

	// Your Gin route setup here
	EmailRoute(router, dbMS, dbMY, validate)

	return router
}

func EmailRoute(router *gin.Engine, dbMS *gorm.DB, dbMY *gorm.DB, validate *validator.Validate) {
	emailService := service.NewEmailService(
		repository.NewEmailRepository(),
		dbMS,
		dbMY,
		validate,
	)
	emailController := controller.NewEmailController(emailService)

	// Define your Gin routes here
	router.GET("/", auth.Auth(emailController.EmailProsess, []string{}))
	// Add more routes as needed
}
