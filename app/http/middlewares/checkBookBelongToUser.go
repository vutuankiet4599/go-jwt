package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vutuankiet4599/go-jwt/app/models"
	"github.com/vutuankiet4599/go-jwt/helper"
	"gorm.io/gorm"
)

func CheckBookBelongToUser(db *gorm.DB) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		userId := helper.GetUserIdFromToken(ctx)
		if userId == 0 {
			response := helper.BuildErrorResponse("Authorization failed", "User not valid", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		id := helper.GetIdFromRouteParams(ctx)
		if id == 0 {
			response := helper.BuildErrorResponse("Authorization failed", "Route param not valid", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var book models.Book
		err := db.First(&book, id)
		if err.Error != nil {
			response := helper.BuildErrorResponse("Authorization failed", err.Error.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if book.UserId != userId {
			response := helper.BuildErrorResponse("Authorization failed", "User do not own this book", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Next()
	}
}
