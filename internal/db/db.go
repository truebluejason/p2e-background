package db

import (
	"database/sql"
	"../syncLog"
	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	Id int
	Type string
	Author string
	Payload string
}

func GetUsersFromTime(formattedTime string) ([]string, error) {
	
}

func GetRandomContent() (Content, error) {

}