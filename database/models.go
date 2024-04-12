package database

import (
	//"github.com/jackc/pgtype"
)

type PublicSomething struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type PrivateSomething struct {
	ID       int                `json:"id"`
	Text     string             `json:"text"`
	Location string             `json:"location"`
	//Time     pgtype.Timestamptz `json:"time"`
}
