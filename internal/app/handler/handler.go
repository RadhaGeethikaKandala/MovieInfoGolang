package handler

import (
	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {

	ctx.String(200, "hello world!")
}
