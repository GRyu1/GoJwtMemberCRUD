package routes

import (
	"github.com/gin-gonic/gin"
	Controllers "goLangJwtPrac/controllers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.POST("/login", Controllers.HandleAuthentication)
	router.POST("/authorization", Controllers.HandleAuthorization)
	router.GET("/users", Controllers.HandleGetAllUser)
	router.GET("/users/:id", Controllers.HandleGetUser)
	router.POST("/users", Controllers.HandlePostInsertUser)
	router.DELETE("/users/:id", Controllers.HandleDeleteUser)
	router.PATCH("/users/:id", Controllers.HandlePatchUser)

	return router
}
