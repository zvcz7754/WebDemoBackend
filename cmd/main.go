package main

import (
	"webDemoBackend/internal/repository"
	"webDemoBackend/internal/router"
)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

	db777 := repository.ConnectMySQL()
	repository.InserData(db777)
}
