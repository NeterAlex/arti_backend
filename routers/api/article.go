package api

import (
	"arti_backend/models"
	"arti_backend/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {
	capacity, _ := strconv.Atoi(c.Query("capacity"))
	id := c.Query("id")
	title := c.Query("title")
	author := c.Query("author")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	if id != "" {
		maps["id"] = id
	}
	if title != "" {
		maps["title"] = title
	}
	if author != "" {
		maps["author"] = author
	}

	data["lists"] = models.GetArticles(capacity, maps)
	data["total"] = models.GetArticlesTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   data,
	})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := e.INVALID_PARAMS
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	title := c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	ptime := c.Request.FormValue("ptime")
	category := c.Request.FormValue("category")
	preview := c.Request.FormValue("preview")
	content := c.Request.FormValue("content")

	code := e.SUCCESS
	//if !models.ExistArticleByTitle(title) {
	if true {
		code = e.SUCCESS
		models.AddArticle(title, author, ptime, category, preview, content)
	} else {
		code = e.ERROR_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	ptime := c.Request.FormValue("ptime")
	category := c.Request.FormValue("category")
	preview := c.Request.FormValue("preview")
	content := c.Request.FormValue("content")
	code := e.SUCCESS
	if models.ExistArticleByID(id) {
		data := make(map[string]interface{})
		if title != "" {
			data["title"] = title
		}
		if author != "" {
			data["author"] = author
		}
		if ptime != "" {
			data["publish_time"] = ptime
		}
		if category != "" {
			data["category"] = category
		}
		if preview != "" {
			data["preview_image"] = preview
		}
		if content != "" {
			data["content"] = content
		}
		for item := range data {
			fmt.Println(item, data[item])
		}
		models.EditArticle(id, data)
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := e.SUCCESS
	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
