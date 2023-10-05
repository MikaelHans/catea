package main

import (
	"github.com/MikaelHans/catea/login-signup/internal/login"
	"github.com/MikaelHans/catea/login-signup/internal/signup"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", func(c *gin.Context) {
		signup.SignUp(c)
	})
	
	r.POST("/login", func(c *gin.Context) {
		login.Login(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}