package routers

import (
	"UnitConverter/pkg/view"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	app.GET("/", view.ShowForm)
	app.POST("/", view.ConvertUnit)

}
