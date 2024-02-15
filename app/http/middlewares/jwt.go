package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vutuankiet4599/go-jwt/app/service"
	"github.com/vutuankiet4599/go-jwt/helper"
)

func AuthorizeJwt(jwtService service.JwtService) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)
		if !token.Valid {
			log.Println(err)
            response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[userId]: ", claims["userId"])
		ctx.Set("userId", claims["userId"])
		ctx.Next()
	}
}
