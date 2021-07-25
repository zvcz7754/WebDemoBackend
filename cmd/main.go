package main

import (
	router "shoppingWebBackend/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
