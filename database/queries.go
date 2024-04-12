package database

import (
	i "api-starter/internal"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbConn struct {
	Pool *pgxpool.Pool
}

func (db *DbConn)GetPublicSomething(c *gin.Context) (PublicSomething, *i.Error) {
	return PublicSomething{}, nil
}

func (db *DbConn) GetPrivateSomething(c *gin.Context) (PrivateSomething, *i.Error) {
	return PrivateSomething{}, nil
}
