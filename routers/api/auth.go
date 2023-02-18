package api

import (
	"arti_backend/models"
	"arti_backend/pkg/e"
	util "arti_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAuth(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	isExist := models.CheckAuth(username, password)
	if username != "" && password != "" {
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddUser(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	email := c.Request.FormValue("email")

	code := e.INVALID_PARAMS
	if username != "" && password != "" && email != "" {
		if !models.ExistUserByUsername(username) {
			code = e.SUCCESS
			models.AddUser(username, password, email)
		} else {
			code = e.ERROR_REGISTER_EXIST_USER
		}
	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	password := c.Query("password")
	email := c.Query("email")

	code := e.INVALID_PARAMS
	if models.ExistArticleByID(id) {
		data := make(map[string]interface{})
		if password != "" {
			data["password"] = password
		}
		if email != "" {
			data["email"] = email
		}
		models.EditUser(id, data)
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	code := e.INVALID_PARAMS
	if models.ExistUserByUsername(username) {
		code = e.SUCCESS
		models.DeleteUser(models.GetUserByUsername(username).ID)
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func GetUserInfoByName(c *gin.Context) {
	username := c.Param("username")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if username != "" {
		code = e.SUCCESS
		auth := models.GetUserByUsername(username)
		data["username"] = auth.Username
		data["email"] = auth.Email
		data["avatar"] = auth.Avatar

	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
