package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	_, err := fmt.Println("hello world")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "result")
	} else {
		ctx.JSON(http.StatusOK, "result")
	}
}
