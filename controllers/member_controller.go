package controllers

import "github.com/gin-gonic/gin"

func IndexDisplayAction(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func UserAddDisplayAction(c *gin.Context) {
	c.HTML(200, "userAdd.html", gin.H{})
}

func UserEditDisplayAction(c *gin.Context) {
	id := c.Param("id")
	c.HTML(200, "userEdit.html", gin.H{"id": id})
}

func UserListDisplayAction(c *gin.Context) {
	c.HTML(200, "userList.html", gin.H{})
}
