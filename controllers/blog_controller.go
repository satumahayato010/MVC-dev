package controllers

import (
	"MVC-dev/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	datas := models.GetAll()
	c.HTML(http.StatusOK, "index.html", gin.H{"datas": datas})
}

func ShowOneBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("No such ID")
	}
	data := models.GetOne(id)
	c.HTML(http.StatusOK, "show.html", gin.H{"data": data})
}

func ShowCreateBlog(c *gin.Context) {
	c.HTML(200, "create.html", nil)
}

func CreateBlog(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	data := models.BlogEntity{Title: title, Body: body}
	data.Create()
	c.Redirect(301, "/")
}

func ShowEditBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("ERROR")
	}
	data := models.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"data": data})
}

func EditBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	data := models.GetOne(id)
	title := c.PostForm("title")
	data.Title = title
	body := c.PostForm("body")
	data.Body = body
	data.Update()
	c.Redirect(301, "/")
}

func ShowCheckDeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"data": data})
}

func DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.GetOne(id)
	data.Delete()
	c.Redirect(301, "/")
}
