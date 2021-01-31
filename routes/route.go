package routes

import (
	control "Uber/controller"
	m "Uber/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", control.Signup)
	router.POST("/login", control.Login)
	router.POST("/logout", m.TokenAuthMiddleware(), control.Logout)
	router.POST("/token/refresh", control.Refresh)
	router.POST("/searchnearcabs", m.TokenAuthMiddleware(), control.NearMyCab)
	router.POST("/bookride", m.TokenAuthMiddleware(), control.BookRide)

	return router
}
