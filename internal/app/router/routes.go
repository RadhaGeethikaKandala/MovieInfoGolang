package router

import (
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/handler"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {

	engine.GET("/hello", handler.SayHello)
}
