package redisRepository

import (
	"errors"
	"fmt"
	"ozon/internal/models"
	"time"
)

func CheckIfExist(fullLink string) bool {
	val := models.Redis.Connection.Exists(fullLink)
	isExist, _ := val.Result()
	if isExist == 1 {
		return true
	}
	return false
}

func CreateNewNode(short string, long string) models.OwnError {
	err := models.Redis.Connection.Set(short, long, 0).Err()
	if err != nil {
		return models.OwnError{Err: errors.New("repository error"), Code: 500, Timestamp: time.Now(), Message: "repository error"}
	}
	return models.OwnError{}
}

func GetFullLink(short string) (string, models.OwnError) {
	result, err := models.Redis.Connection.Get(short).Result()
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return "", models.OwnError{Err: errors.New("repository error"), Code: 500, Timestamp: time.Now(), Message: "repository error"}
	}
	return result, models.OwnError{}
}
