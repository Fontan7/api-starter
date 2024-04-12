package database

import (
	i "api-starter/internal"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	//Get
	GetPublicSomething(c *gin.Context) (PublicSomething, i.Error)
	GetPrivateSomething(c *gin.Context) (PrivateSomething, i.Error)
}

// DatabaseConnectPGX will connect to the database using the pgx driver
func NewDatabasePool(conf i.DatabaseConfig) (*DbConn, error) {
	fmt.Println("Opening database pool...")

	// Connect to the database
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, conf.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	fmt.Println("Health check...")
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to verify connection to database: %v", err)
	}

	// Return the connection
	fmt.Println("Successfully connected to database")
	return &DbConn{Pool: pool}, nil
}
