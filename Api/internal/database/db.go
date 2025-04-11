package database

import (
	"fmt"
	"log"

	"github.com/HugoCBB/api-contrato-storage/internal/config"
	"github.com/HugoCBB/api-contrato-storage/internal/models"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() error {
	err = godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
		return err
	}

	dbConfig := config.NewConfigDatabase()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao se conectar ao banco ", err)
		return err

	}

	fmt.Println("Conexao bem sucedida com o banco de dados")

	DB.AutoMigrate(&models.User{})
	return nil

}
