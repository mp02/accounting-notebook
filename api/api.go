package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetCapital ...
var GetCapital = func(ctx *gin.Context) {

	dato, _ := strconv.Atoi(ctx.Param("id"))
	ctx.JSON(http.StatusOK, dato)

}

//Posting post
var PostCredit = func(ctx *gin.Context) {

	contenttype := ctx.Request.Header.Get("Content-Type")

	ctx.JSON(http.StatusCreated, contenttype)

}

var PostDebit = func(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"Debitado": nil})
}
