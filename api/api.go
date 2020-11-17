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

// GetCapital godoc
// @Summary Return an integer
// @Description The amount is the capital of the user.
// @Accept  json
// @Produce  json
// @Param id path int true "UserID"
// @Success 200 float32 Example
// @Failure 500 {object} error
// @Router /capital [get]
var GetCapital = func(ctx *gin.Context) {

	capital := user.GetCapital()
	ctx.JSON(http.StatusOK, capital)
	return
}

//PostCredit godoc
// @Summary Post credit into an account
// @Description .
// @Accept  json
// @Produce  json
// @Param Request body models.BodyCredit true "Info for the transaction"
// @Success 200 {object} models.User
// @Failure 500 {object} error
// @Router /capital [get]
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

// PostDebit
// @Summary Post debit into an account
// @Description .
// @Accept  json
// @Produce  json
// @Param Request body models.BodyDedit true "Info for the transaction"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} error
// @Router /capital [get]
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
