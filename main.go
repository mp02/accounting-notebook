package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mp02/accounting-notebook/api"
)

func main() {
	router := gin.Default()

	//authorized := router.Group("/")
	//authorized.Use(AuthRequired())

	{
		router.GET("/capital/:id", api.GetCapital)
		router.POST("/credit/:id", api.PostCredit)
		router.POST("/debit/:id", api.PostDebit)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
