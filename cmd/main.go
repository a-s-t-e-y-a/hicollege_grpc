package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"hicollege/config/pg"
	"hicollege/config/redis"
	userpb "hicollege/gen/go/create_user"
	onboardingpb "hicollege/gen/go/on_boarding"
	models "hicollege/migrations"
	"hicollege/service/onboarding"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := pg.DbOptions{
		Host:     ":5432",
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Db:       os.Getenv("DATABASE_NAME"),
	}
	db := dbConfig.Connection()
	defer db.Close()
	if err := models.RegisterModels(db); err != nil {
		log.Fatalf("Failed to register database models: %v", err)
	}
	redisConfig := redis.RedisConnectionOptions{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Protocol: 2,
	}

	redisDb := redisConfig.RedisConnection()
	defer redisDb.Close()

	server := grpc.NewServer()

	sendOtpService := onboarding.NewOtpService(redisDb)
	createUserService := onboarding.NewCreateUserService(db)
	onboardingpb.RegisterOtpServiceServer(server, sendOtpService)
	userpb.RegisterUserServiceServer(server, createUserService)

	lis, err := net.Listen("tcp", ":55053")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	fmt.Println(lis, "Connection started at the endpoint 55053")

	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to serv", err)
	}

}
