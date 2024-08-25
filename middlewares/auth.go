package middlewares

import (
	"net/http"

	"example.com/gin-go-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(cntx *gin.Context) {
	token := cntx.Request.Header.Get("Authorization")

	if token == "" {
		cntx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Not authenticated"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		cntx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Not authenticated", "data": err})
		return
	}

	cntx.Set("userId", userId)

	cntx.Next()
}
