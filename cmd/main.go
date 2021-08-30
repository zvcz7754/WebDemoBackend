package main

import (
	_ "github.com/tedmax100/gin-angular/docs"
	"webDemoBackend/internal/router"
)

// @title go backend demo
// @version 1.0

// @contact.name Yi
// @contact.url https://tedmax100.github.io/

// @license.name Apache 2.0
// @BasePath /api/v1
// @host localhost:8080
// schemes http
func main() {
	//repository.InitDataBase()

	router.InitRouter()

}
