package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"ozon/config"
	"ozon/internal/models"
	"ozon/internal/repository"
	"ozon/internal/restDelivery"
	methods "ozon/internal/server"
	"ozon/pkg/api"
	"ozon/pkg/db"
	testing "ozon/test"
)

func main() {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	models.DB = os.Getenv("DATABASE")
	log.Print(fmt.Sprintf("Storage database - %s", models.DB))
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}
	if models.DB == "redis" {
		redisConnection(conf)
	} else if models.DB == "postgres" {
		postgresConnection(conf)
	} else if models.DB == "all" {
		redisConnection(conf)
		postgresConnection(conf)
	} else {
		log.Fatal("Error with db params")
	}
	service := os.Getenv("SERVICE")
	log.Print(fmt.Sprintf("Sevice - %s", service))
	if service == "grpc" {
		grpcService()
	} else if service == "rest" {
		restService()
	} else if service == "test" {
		testing.Test()
	} else {
		log.Fatal("Error with service params")
	}
}

func grpcService() {
	server := grpc.NewServer()
	srv := &methods.Server{}
	api.RegisterLinkChangerServer(server, srv)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Print(err)
	}
	if err := server.Serve(listener); err != nil {
		log.Print(err)
	}
}

func postgresConnection(conf *config.Config) {
	var err error
	models.Tools.Connection, err = db.InitPsqlDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = models.Tools.Connection.Ping()
	if err != nil {
		log.Panic(err)
	}
	repErr := repository.InitTable()
	if repErr.Err != nil {
		log.Fatal(repErr.Err)
	}
}

func redisConnection(conf *config.Config) {
	models.Redis.Connection = db.InitRedis(conf)

	_, err := models.Redis.Connection.Ping().Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(fmt.Sprintf("Connection Redis success. Host: %s\tPort: %s", *conf.Redis.Host, *conf.Redis.Port))
	}
}

func restService() {
	var app = fiber.New()
	restDelivery.Hearing(app)
}

