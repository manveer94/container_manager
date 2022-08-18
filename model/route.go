package model

import "github.com/gin-gonic/gin"

type Route struct {
	Path     string
	Method   string
	Callback func(context *gin.Context)
}
