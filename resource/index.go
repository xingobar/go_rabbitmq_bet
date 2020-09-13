package resource

import (
	"github.com/gin-gonic/gin"
	"go_rabbitmq_bet/pkg/e"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})
}
