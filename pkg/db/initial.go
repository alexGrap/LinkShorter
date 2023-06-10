package db

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	"ozon/config"
)

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*c.Postgres.Host,
		*c.Postgres.Port,
		*c.Postgres.User,
		*c.Postgres.Password,
		*c.Postgres.DbName)
	fmt.Println(connectionUrl)
	database, err := sqlx.Connect("pgx", connectionUrl)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func InitRedis(c *config.Config) *redis.Client {
	opt := redis.Options{Addr: fmt.Sprintf("%s:%s", *c.Redis.Host, *c.Redis.Port), Password: *c.Redis.Password, DB: *c.Redis.DbName}
	client := redis.NewClient(&opt)
	return client
}
