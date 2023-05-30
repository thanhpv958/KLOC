package main

import (
	"kloc-go/config/db"
	"kloc-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	godotenv.Load()   // init env
	routes.InitAPI(r) // init api route
	db.Init()         // init db

	r.Run()
}
