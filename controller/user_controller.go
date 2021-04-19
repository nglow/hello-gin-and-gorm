package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"helloGinAndGorm/dto"
	"helloGinAndGorm/helper"
	"helloGinAndGorm/service"
	"net/http"
	"strconv"
)

// UserController is a ....
type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JwtService
}

func NewUserController(userService service.UserService, jwtService service.JwtService) UserController {
	return &userController {
		userService: userService,
		jwtService: jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDto dto.UserUpdateDto
	errDto := ctx.ShouldBind(&userUpdateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["UserId"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDto.Id = id
	uRes := c.userService.Update(userUpdateDto)
	res := helper.BuildResponse(true, "OK!", uRes)
	ctx.JSON(http.StatusOK, res)
	return
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["UserId"]))
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
	return
}