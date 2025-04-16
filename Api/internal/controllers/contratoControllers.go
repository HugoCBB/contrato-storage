package controllers

import (
	"net/http"
	"time"

	"github.com/HugoCBB/api-contrato-storage/internal/database"
	"github.com/HugoCBB/api-contrato-storage/internal/models"
	"github.com/gin-gonic/gin"
)

func Contratos(c *gin.Context) {
	//Mostra todos os contratos
	var t []models.Contrato
	database.DB.Find(&t)
	c.JSON(http.StatusOK, t)
}

func AdicionarContrato(c *gin.Context) {
	// Criar novo contrato
	var t models.Contrato

	if err := c.ShouldBindBodyWithJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	t.Data_criacao = time.Now().Format("01/02/2006")
	database.DB.Create(&t)

	c.JSON(http.StatusOK, &t)

}

func DeletarContrato(c *gin.Context) {
	// O nome da funcao esta bem explicativa nao acha?
	var t models.Contrato
	id := c.Param("id")
	database.DB.Delete(&t, id)
	c.JSON(http.StatusOK, gin.H{"status": "contrato deletado com sucesso"})

}
