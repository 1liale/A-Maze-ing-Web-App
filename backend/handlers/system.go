package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SystemCheck(ctx *gin.Context) {
	obj := make(map[string]interface{})
	obj["status"] = 200
	obj["health"] = "GOOD"

	ctx.JSON(http.StatusOK, obj)
}
