package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mp02/accounting-notebook/api"
)

// @title Accounting-notebook
// @version 1.0
// @description API Restful for credit and debit transactions

// @contact.name Martin Pruyas
// @contact.url https://www.linkedin.com/in/martin-pruyas/
// @contact.email o.gema.pg@gmail.com

// @license.name mp.02

// @host localhost
// @BasePath /v1

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/capital/:id", api.GetCapital)
		v1.POST("/credit/:id", api.PostCredit)
		v1.POST("/debit/:id", api.PostDebit)
	}
	router.Run()
}
