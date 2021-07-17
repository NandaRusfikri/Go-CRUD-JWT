package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"../models"

	"../structs"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	params := structs.ParamsGetListUser{}

	response := structs.Response{}
	users := []structs.DataGetUser{}

	err := c.ShouldBind(&params)
	var Count int

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = "validation " + mess
		c.JSON(400, response)
	} else {
		users, Count, err = models.GetUsers(params)

		if err == nil {
			response.Count = Count
			response.Data = users

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func CreateUser(c *gin.Context) {

	params := structs.ParamsCreateUser{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.CreateUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func UpdateUser(c *gin.Context) {

	params := structs.ParamsUpdateUser{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.UpdateUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func DeleteUser(c *gin.Context) {

	params := structs.ParamsUser{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.DeleteUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func Login(c *gin.Context) {

	params := structs.ParamsLogin{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.Login(params)

		if err == nil {

			if len(Data) > 0 {
				sign := jwt.New(jwt.GetSigningMethod("HS256"))
				token, err := sign.SignedString([]byte("nandarusfikri"))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": err.Error(),
					})
					c.Abort()
				}

				response.Message = token
				response.Data = Data

				c.JSON(http.StatusOK, response)
			} else {
				response.Message = "Username / Password Salah"

				c.JSON(http.StatusOK, response)
			}

		} else {
			c.JSON(500, response)
		}
	}

}
