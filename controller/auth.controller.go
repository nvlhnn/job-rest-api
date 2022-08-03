package controller

import (
	"dansmultipro/recruitment/dto"
	"dansmultipro/recruitment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) 
}


type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}


func LoginHandler(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}


func (c *loginController) Login(ctx *gin.Context)  {

	var credential dto.LoginDTO
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	isUserAuthenticated, res := c.loginService.VerifyCredential(credential.Username, credential.Password)
	if isUserAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{
			"token": c.jWtService.GenerateToken(res),
		})
	}else{
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	
}

