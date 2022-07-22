package main

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middleware(ctx *gin.Context) {
	var userType string
	if val, exists := ctx.Get("user-type"); exists {
		userType = val.(string)
	}

	if userType == "" {
		userType = "gen-user"
	}

	tokenBucket := GetBucket(GetClientIndentifire(ctx), userType)

	if !tokenBucket.IsRequesrAllowed(1) {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Try again after sometime!",
		})
		return
	}
	ctx.Next()
}

func GetClientIndentifire(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	url := ctx.Request.URL.Path
	data := fmt.Sprintf("%s-%s", ip, url)
	h := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(h[:]).Text(62)
	return hash
}
