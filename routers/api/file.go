package api

import (
	"arti_backend/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadAvatar(c *gin.Context) {
	code := e.INVALID_PARAMS
	//username := c.Request.FormValue("username")
	file, err := c.FormFile("image")
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	log.Println("An avatar image was uploaded: " + file.Filename)
	dst := fmt.Sprintf("./database/avatar/%s", file.Filename)
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})

}
