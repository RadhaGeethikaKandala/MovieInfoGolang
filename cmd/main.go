package main

import (
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Router(engine)

	engine.Run(":8085")
}
