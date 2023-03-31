package app

import (
	"fmt"
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

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.New()
	router.Use(ErrorHandler())
	route.EmailRoute(router, db, validate)

	return router
}
