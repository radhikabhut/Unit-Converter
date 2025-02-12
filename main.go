package main

import (
	"UnitConverter/pkg/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.Static("/static", "./static")
	app.LoadHTMLGlob("templates\\*")
	
	routers.InitRouter(app)

	app.Run(":8080")

}
