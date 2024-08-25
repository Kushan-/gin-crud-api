package routes

import (
	"fmt"
	"net/http"

	"example.com/gin-go-api/models"
	"example.com/gin-go-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(cntx *gin.Context) {
	var user models.User
	err := cntx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("Signup user -> err", err)
		fmt.Println()
		cntx.JSON(http.StatusBadRequest, gin.H{"msg": "sign up err", "data": err})
		return
	}

	err = user.SaveToQL()

	if err != nil {
		fmt.Println("SAVE to QL USER -> err", err)
		fmt.Println()
		cntx.JSON(http.StatusBadRequest, gin.H{"msg": "sign up err", "data": err})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{"mag": "User created"})
}

func login(cntx *gin.Context) {
	var user models.User
	err := cntx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("Signup user -> err", err)
		fmt.Println()
		cntx.JSON(http.StatusBadRequest, gin.H{"msg": "sign up err", "data": err})
		return
	}

	err = user.ValidateCreds()

	if err != nil {
		cntx.JSON(http.StatusUnauthorized, gin.H{"msg": "sign up err", "data": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		cntx.JSON(http.StatusUnauthorized, gin.H{"msg": "Couldnt authenticate", "data": err})
		return

	}
	fmt.Println(token)
	cntx.JSON(http.StatusOK, gin.H{"mag": "Login Successful!!!", "token": token})

}
