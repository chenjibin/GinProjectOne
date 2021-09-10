package main

import (
	"GinProjectOne/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	userRoutes := r.Group("/user")
	userRoutes.GET("/info", controller.GetUserInfo)
	userRoutes.GET("/list", controller.GetUserList)
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	return r
}
