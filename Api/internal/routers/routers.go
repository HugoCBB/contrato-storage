package routers

import (
	"github.com/HugoCBB/api-contrato-storage/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/", controllers.AllUser)
			user.POST("/register", controllers.AdicionarUsuario)
		}
	}

	r.Run()
}
