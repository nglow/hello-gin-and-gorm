package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"helloGinAndGorm/dto"
	"helloGinAndGorm/entity"
	"helloGinAndGorm/helper"
	"helloGinAndGorm/service"
	http "net/http"
	"strconv"
)

// BookController is a ....
type BookController interface {
	All(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService service.JwtService
}

// NewBookController creates a new instances of BookController
func NewBookController(bookService service.BookService, jwtService service.JwtService) BookController {
	return &bookController{
		bookService: bookService,
		jwtService: jwtService,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()
	res := helper.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book entity.Book = c.bookService.FindById(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("Data not found,", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDto dto.BookCreateDto
	errDto := context.ShouldBind(&bookCreateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userId := c.getUserIdByToken(authHeader)
		convertUserId, err := strconv.ParseUint(userId, 10, 64)
		if err == nil {
			bookCreateDto.UserId = convertUserId
		}
		result := c.bookService.Insert(bookCreateDto)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDto dto.BookUpdateDto
	errDto := context.ShouldBind(&bookUpdateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["UserId"])
	if c.bookService.IsAllowedToEdit(userId, bookUpdateDto.Id) {
		id, errId := strconv.ParseUint(userId, 10, 64)
		if errId == nil {
			bookUpdateDto.UserId = id
		}
		result := c.bookService.Update(bookUpdateDto)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) Delete(ctx *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	book.Id = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["UserId"])
	if c.bookService.IsAllowedToEdit(userId, book.Id) {
		c.bookService.Delete(book)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) getUserIdByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["UserId"])
}
