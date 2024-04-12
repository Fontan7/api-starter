package service

import (
	"api-starter/database"
	"api-starter/internal"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt"
)

func setDb(db *database.DbConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func setApp(app internal.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	}
}

func corsConfig(*gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTION")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-API-Key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		//fmt.Println(c.Request.Method)
		//fmt.Println(c.Request.Response)
		//fmt.Println(c.Request.WithContext(c))
	}
}

func checkAPIKey(clientKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if clientKey != c.GetHeader("X-API-Key") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid api key")
		}

		c.Next()
	}
}

func validateAndSetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//type TokenClaims struct {
		//	Email string `json:"email"`
		//	jwt.StandardClaims
		//}

		accessToken := c.GetHeader("Authorization")
		//accessToken = accessToken[len("bearer "):]

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: missing token")
		}
		fmt.Println(accessToken)

		/*
			claims := &TokenClaims{}
			token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				// Make sure that the token method conforms to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("ValidateToken unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(internal.AppConfig.JwtKey()), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: invalid signature: "+err.Error())
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, "validateToken error: "+err.Error())
			}
			if !token.Valid {
				c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: invalid token: "+err.Error())
			}

			//do claims validation here


			c.Set("token", token)
			c.Set("claims", claims)
		*/
		c.Next()
	}
}

type ControllerFunctions func(c *gin.Context) (response interface{}, err *internal.Error)

func gHandler(c *gin.Context, fn ControllerFunctions) {
	response, err := fn(c)
	if err != nil {
		path := c.Request.URL.Path
		err.Path = path
		c.AbortWithStatusJSON(err.Status, err)
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
