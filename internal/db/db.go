package db

import (
	"github.com/truebluejason/p2e-background/internal/conf"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Content struct {
	Id int
	Type string
	Author string
	Payload string
}

var dataSource string = conf.Configs.DBUser + ":" + conf.Configs.DBPassword + "@/" + conf.Configs.DBName

func GetUsersFromTime(formattedTime string) ([]string, error) {
	var userIDs []string

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
	    return userIDs, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
	    return userIDs, err
	}

	rows, err := db.Query("SELECT UserID FROM UserTimes WHERE Stamp = ?", formattedTime)
	if err != nil {
		return userIDs, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			return userIDs, err
		}
		userIDs = append(userIDs, userID)
	}

	err = rows.Err()
	return userIDs, err
}

func GetRandomContent() (Content, error) {
	var randContent Content = Content{}
	var author sql.NullString

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
	    return randContent, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
	    return randContent, err
	}

	rows, err := db.Query("SELECT ContentID, Type, Author, Content FROM Contents ORDER BY Rand() LIMIT 1")
	if err != nil {
		return randContent, err
	}
	defer rows.Close()

	rows.Next()
	if err := rows.Scan(&randContent.Id, &randContent.Type, &author, &randContent.Payload); err != nil {
		return randContent, err
	}

	if author.Valid {
		randContent.Author = author.String
	} else {
		randContent.Author = ""
	}
	err = rows.Err()
	return randContent, err
}
