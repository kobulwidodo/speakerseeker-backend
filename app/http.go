package main

import (
	"fmt"
	"log"
	"os"
	"speakerseeker-backend/domain"
	_speakerHttpHandler "speakerseeker-backend/speaker/delivery/http"
	_speakerRepository "speakerseeker-backend/speaker/repository/postgresql"
	_speakerUsecase "speakerseeker-backend/speaker/usecase"
	_speakerSkillHttpHandler "speakerseeker-backend/speaker_skill/delivery/http"
	_speakerSkillRepository "speakerseeker-backend/speaker_skill/repository/postgresql"
	_speakerSkillUsecase "speakerseeker-backend/speaker_skill/usecase"
	_userHttpHandler "speakerseeker-backend/user/delivery/http"
	_userRepository "speakerseeker-backend/user/repository"
	_userUsecase "speakerseeker-backend/user/usecase"

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
	speakerRepository := _speakerRepository.NewSpeakerRepository(db)
	speakerSkillRepository := _speakerSkillRepository.NewSpeakerSkillRepository(db)

	userUseCase := _userUsecase.NewUserUseCase(userRepository)
	speakerUseCase := _speakerUsecase.NewSpeakserUsecase(speakerRepository)
	speakerSkillUseCase := _speakerSkillUsecase.NewSpeakerSkillUsecase(speakerSkillRepository)

	_userHttpHandler.NewUserHandler(api, userUseCase)
	_speakerHttpHandler.NewSpeakerHandler(api, speakerUseCase)
	_speakerSkillHttpHandler.NewSpeakerSkillHandler(api, speakerSkillUseCase)

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

	if err := db.AutoMigrate(&domain.User{}, &domain.Speaker{}, &domain.SpeakerSkill{}, &domain.SpeakerCareer{}); err != nil {
		return nil, err
	}

	return db, nil
}
