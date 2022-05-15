package main

import (
	"fmt"
	"speakerseeker-backend/domain"
	_userHttpHandler "speakerseeker-backend/user/delivery/http"
	_userRepository "speakerseeker-backend/user/repository"
	_userUsecase "speakerseeker-backend/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
		panic(err)
	}

	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	r := gin.Default()
	api := r.Group("/api")

	userRepository := _userRepository.NewUserRepository(db)

	userUseCase := _userUsecase.NewUserUseCase(userRepository)

	_userHttpHandler.NewUserHandler(api, userUseCase)

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
