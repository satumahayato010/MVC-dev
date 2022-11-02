package main

import "MVC-dev/controllers"

func main() {
	router := controllers.GetRouter()
	router.Run("localhost:8080")
}
