package main

import (
	"de.anikate/blog-api/db"
	"de.anikate/blog-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	engine := gin.Default()

	routes.Setup(engine)

	engine.Run(":8080")
}
