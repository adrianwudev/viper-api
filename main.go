package main

import (
	"example.com/viper/repository"
	"example.com/viper/routes"
)

func main() {
	repository.New()
	routes.Start()
}
