package controller

import db "api-starter/database"

// Response is a struct that contains the response data
type ResponsePublic struct {
	db.PublicSomething
}

type ResponsePrivate struct {
	db.PrivateSomething
}
