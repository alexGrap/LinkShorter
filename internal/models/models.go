package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB string

type OwnError struct {
	Code      int
	Err       error
	Message   string
	Timestamp time.Time
}

type Link struct {
	shortLink string
	longLink  string
}

func (err *OwnError) Error() string {
	return fmt.Sprintf("Status: %d\n"+
		"Error: %s\n"+
		"Message: %s\n"+
		"Timestamp: %s\n",
		err.Code, err.Err.Error(), err.Message, err.Timestamp.String())
}

func (err *OwnError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code      int    `json:"code"`
		Err       string `json:"err"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
	}{
		Code:      err.Code,
		Err:       err.Err.Error(),
		Message:   err.Message,
		Timestamp: err.Timestamp.String(),
	})
}

var Tools tools

type tools struct {
	Connection *sqlx.DB
}

var Redis redisTool

type redisTool struct {
	Connection *redis.Client
}

type RestResult struct {
	ResultLink string `json:"resultLink"`
}
