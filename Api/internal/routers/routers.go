package routers

import (
	"github.com/HugoCBB/api-contrato-storage/internal/controllers"
	"github.com/HugoCBB/api-contrato-storage/internal/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.ContentTypeMiddleware())

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/", controllers.Usuarios)
			user.POST("/register", controllers.AdicionarUsuario)
		}

		contrato := api.Group("/contrato")
		{
			contrato.GET("/", controllers.Contratos)
			contrato.POST("/register", controllers.AdicionarContrato)
			contrato.DELETE("/deletar/:id", controllers.DeletarContrato)
		}
	}

	r.Run(":8080")
}
