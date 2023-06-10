package repository

import (
	"errors"
	"ozon/internal/models"
	"time"
)

func CheckIfExist(fullLink string) (string, error) {
	var result string
	err := models.Tools.Connection.QueryRowx(`SELECT shortlink FROM links WHERE longlink = $1`, fullLink).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func CreateNewNode(short string, long string) models.OwnError {
	_, err := models.Tools.Connection.Exec(`INSERT INTO links (shortlink, longlink)VALUES ($1, $2)`, short, long)
	if err != nil {
		return models.OwnError{Err: errors.New("repository error"), Code: 500, Timestamp: time.Now(), Message: "repository error"}
	}
	return models.OwnError{}
}

func GetFullLink(short string) (string, models.OwnError) {
	var result string
	err := models.Tools.Connection.QueryRowx(`SELECT longlink FROM links WHERE shortlink = $1`, short).Scan(&result)
	if err != nil {
		return "", models.OwnError{Err: errors.New("repository error"), Code: 500, Timestamp: time.Now(), Message: "repository error"}
	}
	return result, models.OwnError{}
}

func InitTable() models.OwnError {
	_, err := models.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  links (
    	shortlink TEXT PRIMARY KEY NOT NULL ,
    	longlink TEXT NOT NULL
		);`)
	if err != nil {
		return models.OwnError{Err: errors.New("creation repository error"), Message: "repository error"}
	}
	return models.OwnError{}
}
