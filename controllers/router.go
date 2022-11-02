package controllers

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", IndexDisplayAction)
	r.GET("/user", UserAddDisplayAction)
	r.GET("/user/:id", UserEditDisplayAction)
	r.GET("/user/edit/:id", UserListDisplayAction)

	return r
}
