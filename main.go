package main

import (
	"ans/controllers"
)

func main() {
	libController := controllers.NewController()
	libController.Run()
}
