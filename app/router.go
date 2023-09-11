package app

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"auto-emails/exception"
	"auto-emails/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err)
			}
		}()
		c.Next()
	}
}

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Endpoint not found",
		})
	}
}

func NewRouter(dbMS *gorm.DB, dbMY *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.New()
	router.Use(ErrorHandler())
	router.NoRoute(NotFoundHandler())
	route.EmailRoute(router, dbMS, dbMY, validate)

	return router
}
