package main

import (
	"log"
	"simplefast/core"
	"simplefast/middlewares"
)

func Testing(c *core.Context) error {

	log.Println("sdfsdfsd")
	log.Println(c.Query("tes"))

	return core.ResponseText(c, 200, "Welcome to the Home Page!")
}

func Testing2(c *core.Context) error {

	log.Println("sdfsdfsd")
	log.Println(c.Query("tes"))
	return core.ResponseSuccess(c, "sss")
}

func TestPost(c *core.Context) error {

	log.Println("sdfsdfsd")
	log.Println(c.File("tes"))
	log.Println(c.JSONBody("aka"))
	log.Println("c.Body")
	log.Println(c.Body("tes"))
	return core.ResponseSuccess(c, "sss")
}

func main() {

	router := core.NewRouter()
	router.GET("/home", Testing)
	router.GET("/home2", Testing2)

	router.POST("/post", TestPost)
	router.PUT("/put", TestPost)
	router.DELETE("/delete", TestPost)

	router.Use("/home", middlewares.LoggingMiddleware)
	router.Use("/home", middlewares.LoggingMiddlewareFaile)

	core.RunServer(router, 8080)
}
