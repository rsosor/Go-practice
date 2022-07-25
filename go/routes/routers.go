package routes

import (
	"go-hello/go/controller"

	_ "go-hello/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": "404", "message": "Page not found",
		})
	})

	router.Group("/demo/v1", gin.BasicAuth(gin.Accounts{
		"user1": "love",
		"user2": "god",
		"user3": "sex",
	})).
		GET("/GetUserList", controller.GetUserList).
		POST("/CreateUser", controller.CreateUser).
		PUT("/users/:id", controller.UpdateUser).
		DELETE("/users/:id", controller.DeleteUserById).
		GET("/users/:id", controller.GetUserById)
		// POST("/postName", controller.postName).
		// POST("/updateName", controller.updateName).
		// POST("/deleteName", controller.deleteName)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
