package middleware

import (
	"casbinginxorm/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtAuthorize() gin.HandlerFunc {

	return func(c *gin.Context) {

		// authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "error header"})
			c.Abort()
			return
		}
		code := utils.SUCCESS
		tokenString = tokenString[7:]
		utilsjwt := utils.Jwt{}
		claims, err := utilsjwt.ParserToken(tokenString)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorMalformed:
				code = utils.ValidationErrorMalformed
			case jwt.ValidationErrorExpired:
				code = utils.ValidationErrorExpired
			case jwt.ValidationErrorNotValidYet:
				code = utils.ValidationErrorNotValidYet
			case jwt.ValidationErrorAudience:
				code = utils.ValidationErrorAudience
			case jwt.ValidationErrorIssuer:
				code = utils.ValidationErrorIssuer
			default:
				code = utils.ValidationError
			}
		}

		if code != utils.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  utils.GetMsg(code),
				"data": claims,
			})

			c.Abort()

			return
		}

		c.Next()
	}
}
