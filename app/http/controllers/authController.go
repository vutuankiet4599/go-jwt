package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vutuankiet4599/go-jwt/app/models"
	"github.com/vutuankiet4599/go-jwt/app/request"
	"github.com/vutuankiet4599/go-jwt/app/service"
	"github.com/vutuankiet4599/go-jwt/helper"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login (ctx *gin.Context)
	User (ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JwtService
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerRequest request.RegisterRequest
	errRequest := ctx.ShouldBind(&registerRequest)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerRequest.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email address", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	hashedPassword, err := helper.GenerateHashValue(registerRequest.Password)
	if err != nil {
		response := helper.BuildErrorResponse("Something went wrong", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	registerRequest.Password = hashedPassword

	createdUser, isError, errors := c.authService.CreateUser(registerRequest)

	if isError {
		response := helper.BuildErrorResponse("Failed to process request", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))

	data := &struct {
		models.User
		Token string `json:"token"`
	}{
		*createdUser,
		token,
	}

	response := helper.BuildResponse(true, "OK!", data)

	ctx.JSON(http.StatusCreated, response)
}

func (c *authController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	errRequest := ctx.ShouldBind(&loginRequest)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, isErr, errs := c.authService.VerifyCredentials(loginRequest)
	if isErr {
		response := helper.BuildErrorResponse("Authentication failed", errs, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))

	data := &struct {
		models.User
		Token string `json:"token"`
	}{
		*user,
		token,
	}

	response := helper.BuildResponse(true, "OK!", data)

	ctx.JSON(http.StatusOK, response)
}

func (c *authController) User(ctx *gin.Context) {
	userId := helper.GetUserIdFromToken(ctx)
	if userId == 0 {
		response := helper.BuildErrorResponse("Authorization failed", "User not valid", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, isErr, errs := c.authService.GetCurrentUser(userId)

	if isErr {
		response := helper.BuildErrorResponse("Failed to process request", errs, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)
}
