package main

import (
	"github.com/HugoCBB/api-contrato-storage/internal/database"
	"github.com/HugoCBB/api-contrato-storage/internal/routers"
)

func main() {
	database.ConnectDatabase()
	routers.HandleRequests()
}
