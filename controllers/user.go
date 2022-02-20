package controllers

import (
	"net/http"
	"yesiamdonation/response"
	"yesiamdonation/services"
	"yesiamdonation/services/request"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JwtService
}

func NewAuthController(authService services.AuthService, jwtService services.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var register request.RegisterRequest
	err := ctx.ShouldBind(&register)
	if err != nil {
		response := response.BuildErrorResponse("failed to process request", err.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(register.Email) {
		response := response.BuildErrorResponse("Failed to process request", "Duplicate Email", response.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.Register(register)
		token := c.jwtService.GenerateToken(createUser.ID)
		createUser.Token = token
		response := response.BuildSuccessResponse(true, "OK!", &createUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
