package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromToken(ctx *gin.Context) uint {
	userId := ctx.Keys["userId"]
	if userId == "" {
		return 0
	}
	state, ok := userId.(string)
	if !ok {
		return 0
	}
	validUserId, err := strconv.ParseUint(state, 10, 64)
	if err != nil {
		return 0
	}
	return uint(validUserId)
}

func GetIdFromRouteParams(ctx *gin.Context) uint {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}
