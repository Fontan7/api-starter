package controller

import (
	"net/http"

	"api-starter/internal"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
//	@Summary		Health check
//	@Description	always returns OK
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	string
//	@Success		200
//	@Failure		500
//	@Router			/health [get]
func HealthCheck(c *gin.Context) (interface{}, *internal.Error) {
	return http.StatusOK, nil
}

// ShowAccount godoc
//	@Summary		Get public something
//	@Description	get something by ID
//	@Param			X-API-Key	header	string	true	"secret key"
//	@Produce		json
//	@Param			id	path		int	true	"Something ID"
//	@Success		200	{object}	ResponsePublic
//	@Failure		400	{object}	internal.Error
//	@Failure		404	{object}	internal.Error
//	@Failure		500	{object}	internal.Error
//	@Router			/api-starter/v1/{id} [get]
func GetPublicSomething(c *gin.Context) (interface{}, *internal.Error) {
	return ResponsePublic{}, nil
}

// ShowAccount godoc
//
//	@Summary		Get something
//	@Description	get private something
//	@Param			Authorization	header	string	true	"bearer token"
//	@Param			X-API-Key		header	string	true	"secret key"
//	@Produce		json
//	@Success		200	{object}	ResponsePrivate
//	@Failure		400	{object}	internal.Error
//	@Failure		404	{object}	internal.Error
//	@Failure		500	{object}	internal.Error
//	@Security		ApiKeyAuth
//	@Security		OAuth2Application[read, user]
//	@Router			/api-starter/v1/private/something [get]
func GetPrivateSomething(c *gin.Context) (interface{}, *internal.Error) {
	return ResponsePrivate{}, nil
}