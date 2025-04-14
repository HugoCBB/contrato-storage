package controllers

import (
	"net/http"

	"github.com/HugoCBB/api-contrato-storage/internal/database"
	"github.com/HugoCBB/api-contrato-storage/internal/models"
	"github.com/gin-gonic/gin"
)

func Usuarios(c *gin.Context) {
	// Mostra todos os usuarios
	var u []models.User
	database.DB.Find(&u)
	c.JSON(http.StatusOK, u)

}

func AdicionarUsuario(c *gin.Context) {
	// Cria um novo usuario
	var u models.User

	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&u)
	c.JSON(http.StatusOK, u)
}
