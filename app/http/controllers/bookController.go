package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vutuankiet4599/go-jwt/app/request"
	"github.com/vutuankiet4599/go-jwt/app/service"
	"github.com/vutuankiet4599/go-jwt/helper"
)

type bookController struct {
	bookService service.BookService
}

type BookController interface {
	GetAll(ctx *gin.Context)
	GetOneById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteOneById(ctx *gin.Context)
	DeleteAll(ctx *gin.Context)
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c *bookController) GetAll(ctx *gin.Context) {
	books, isError, errors := c.bookService.GetAll()
	if isError {
		response := helper.BuildErrorResponse("Failed to process request", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.BuildResponse(true, "OK!", books)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) GetOneById(ctx *gin.Context) {
	id := helper.GetIdFromRouteParams(ctx)
	if id == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid route param", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	book, isError, errors := c.bookService.GetOneById(id)
	if isError {
		response := helper.BuildErrorResponse("Failed to process request", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.BuildResponse(true, "OK!", book)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) Insert(ctx *gin.Context) {
	var request request.InsertBookRequest
	errRequest := ctx.ShouldBind(&request)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := helper.GetUserIdFromToken(ctx)
	if userId == 0 {
		response := helper.BuildErrorResponse("Authorization failed", "User not valid", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	data, isError, errors := c.bookService.Insert(&request, userId)
	if isError {
		response := helper.BuildErrorResponse("Authorization failed", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", data)
	ctx.JSON(http.StatusCreated, response)
}

func (c *bookController) Update(ctx *gin.Context) {
	var request request.UpdateBookRequest
	errRequest := ctx.ShouldBind(&request)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := helper.GetIdFromRouteParams(ctx)
	if id == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid route param", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	data, isError, errors := c.bookService.Update(&request, id)
	if isError {
		response := helper.BuildErrorResponse("Authorization failed", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", data)
	ctx.JSON(http.StatusCreated, response)
}

func (c *bookController) DeleteOneById(ctx *gin.Context) {
	id := helper.GetIdFromRouteParams(ctx)
	if id == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid route param", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isError, errors := c.bookService.DeleteOneById(id)
	if isError {
		response := helper.BuildErrorResponse("Authorization failed", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) DeleteAll(ctx *gin.Context) {
	userId := helper.GetUserIdFromToken(ctx)
	if userId == 0 {
		response := helper.BuildErrorResponse("Authorization failed", "User not valid", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	isError, errors := c.bookService.DeleteAll(userId)
	if isError {
		response := helper.BuildErrorResponse("Authorization failed", errors, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
