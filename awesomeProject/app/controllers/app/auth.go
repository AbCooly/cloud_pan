package app

import (
	"awesomeProject/app/common/request"
	"awesomeProject/app/common/response"
	"awesomeProject/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		log.Printf("token creating")
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			log.Printf("success in token err")
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}
