package entity

import "time"

//go:generate easytags $GOFILE db,json

type DBUser struct {
	Id      int        `db:"id" json:"id"`
	Name    string     `db:"name" json:"name"`
	Status  string     `db:"status" json:"status"`
	Created *time.Time `db:"created" json:"created"`
}
