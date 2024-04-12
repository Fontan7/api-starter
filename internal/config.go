package internal

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// ========================================================= Variables =========================================================
var (
	_   = godotenv.Load(".env")
	Env = os.Getenv("ENV")
)

// ========================================================= App ===============================================================
type App interface {
	Env() string
	Port() string
	SupaKey() string
	SupaProjectReference() string
	ClientKey() string
	JwtKey() string
	Host() string
}

type app struct {
	env                  string
	port                 string
	supaKey              string
	supaProjectReference string
	clientKey            string
	jwtKey               string
	host                 string
}

func NewApp() (App, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("ENV variable is empty")
	}
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT variable is empty")
	}

	supaProjectReference := os.Getenv("SUPA_PROJECT_REFERENCE")
	if supaProjectReference == "" {
		return nil, fmt.Errorf("SUPA_PROJECT_REFERENCE variable is empty")
	}

	supaKey := os.Getenv("SUPA_ANON_KEY")
	if supaKey == "" {
		return nil, fmt.Errorf("SUPA_ANON_KEY variable is empty")
	}

	clientKey := os.Getenv("CLIENT_KEY")
	if clientKey == "" {
		return nil, fmt.Errorf("CLIENT_KEY variable is empty")
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return nil, fmt.Errorf("JWT_KEY variable is empty")
	}

	host := os.Getenv("SERVER_HOST")
	if host == "" {
		return nil, fmt.Errorf("SERVER_HOST variable is empty")
	}

	fmt.Println("Successfully loaded app environment variables")
	return &app{
		env:                  env,
		port:                 port,
		supaKey:              supaKey,
		supaProjectReference: supaProjectReference,
		clientKey:            clientKey,
		jwtKey:               jwtKey,
		host:                 host,
	}, nil
}

func (a *app) Env() string                  { return a.env }
func (a *app) Port() string                 { return a.port }
func (a *app) SupaKey() string              { return a.supaKey }
func (a *app) SupaProjectReference() string { return a.supaProjectReference }
func (a *app) ClientKey() string            { return a.clientKey }
func (a *app) JwtKey() string               { return a.jwtKey }
func (a *app) Host() string                 { return a.host }

// ========================================================= DatabaseConfig =========================================================
type DatabaseConfig interface {
	User() string
	Pass() string
	Sock() string
	Host() string
	Name() string
	ConnString() string
}

type database struct {
	user       string
	pass       string
	sock       string
	host       string
	name       string
	connString string
}

func NewDatabaseConfig() (DatabaseConfig, error) {
	/*
	   user := os.Getenv("DB_USER")
	   pass := os.Getenv("DB_PASS")
	   sock := os.Getenv("DB_SOCK")
	   host := os.Getenv("DB_HOST")``
	   name := os.Getenv("DB_NAME")
	*/
	connString := os.Getenv("DB_CONN_STRING")

	if connString == "" {
		return nil, errors.New("new database: one or more environment variables are empty")
	}

	fmt.Println("Successfully loaded database environment variables")
	return &database{
		connString: connString,
	}, nil
}

func (d *database) User() string       { return d.user }
func (d *database) Pass() string       { return d.pass }
func (d *database) Sock() string       { return d.sock }
func (d *database) Host() string       { return d.host }
func (d *database) Name() string       { return d.name }
func (d *database) ConnString() string { return d.connString }

type HandlerResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// ========================================================= Error =========================================================
// Error represents the data returned by an aborted HTTP request.
type Error struct {
	Time   string `json:"time_RFC3339"`
	Path   string `json:"path"`
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// DetailError returns an Error struct with the given status and message.
func DetailError(status int, err error) *Error {
	return &Error{
		Time:   time.Now().Format(time.RFC3339),
		Path:   "",
		Status: status,
		Error:  err.Error(),
	}
}
