package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mp02/accounting-notebook/models"
)

var user models.User

func init() {
	user.UserID = 1234
	user.PersonalInfo.FName = "Homero"
	user.PersonalInfo.LName = "Simpson"
}

//GetCapital ...
var GetCapital = func(ctx *gin.Context) {

	capital := user.GetCapital()
	ctx.JSON(http.StatusOK, capital)

}

//Posting post
var PostCredit = func(ctx *gin.Context) {
	var body models.BodyCredit
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	user.Credit(body.Amount)
	ctx.JSON(http.StatusOK, user)

}

var PostDebit = func(ctx *gin.Context) {
	var body models.BodyDebit
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err = user.Debit(body.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Debit": body.Amount, "Capital": user.Capital})
}
