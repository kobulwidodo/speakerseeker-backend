package main

import (
	"fmt"
	"log"
	"os"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/infrastructure"
	"speakerseeker-backend/middleware"
	_midtransTransactionRepository "speakerseeker-backend/midtrans_transaction/repository/postgresql"
	_speakerHttpHandler "speakerseeker-backend/speaker/delivery/http"
	_speakerRepository "speakerseeker-backend/speaker/repository/postgresql"
	_speakerUsecase "speakerseeker-backend/speaker/usecase"
	_speakerSkillHttpHandler "speakerseeker-backend/speaker_skill/delivery/http"
	_speakerSkillRepository "speakerseeker-backend/speaker_skill/repository/postgresql"
	_speakerSkillUsecase "speakerseeker-backend/speaker_skill/usecase"
	_transactionHttpHandler "speakerseeker-backend/transaction/delivery/http"
	_transactionRepository "speakerseeker-backend/transaction/repository/postgresql"
	_transactionUsecase "speakerseeker-backend/transaction/usecase"
	_userHttpHandler "speakerseeker-backend/user/delivery/http"
	_userRepository "speakerseeker-backend/user/repository"
	_userUsecase "speakerseeker-backend/user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env from local file")
	}

	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	api := r.Group("/api")
	api.Static("/images", "./images")

	jwtMiddleware := middleware.NewAuthMiddleware()
	midtransDriver := infrastructure.NewMidtransDriver()

	userRepository := _userRepository.NewUserRepository(db)
	speakerRepository := _speakerRepository.NewSpeakerRepository(db)
	speakerSkillRepository := _speakerSkillRepository.NewSpeakerSkillRepository(db)
	transactionRepository := _transactionRepository.NewTransactionRepository(db)
	midtransTransactionRepository := _midtransTransactionRepository.NewMidtransTransactionRepository(db)

	userUseCase := _userUsecase.NewUserUseCase(userRepository)
	speakerUseCase := _speakerUsecase.NewSpeakserUsecase(speakerRepository, speakerSkillRepository)
	speakerSkillUseCase := _speakerSkillUsecase.NewSpeakerSkillUsecase(speakerSkillRepository)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(transactionRepository, midtransTransactionRepository, speakerRepository, &midtransDriver)

	_userHttpHandler.NewUserHandler(api, userUseCase, jwtMiddleware)
	_speakerHttpHandler.NewSpeakerHandler(api, speakerUseCase)
	_speakerSkillHttpHandler.NewSpeakerSkillHandler(api, speakerSkillUseCase)
	_transactionHttpHandler.NewTransactionHandler(api, transactionUsecase, jwtMiddleware)

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Speaker{}, &domain.SpeakerSkill{}, &domain.SpeakerCareer{}, &domain.SpeakerExperience{}, &domain.Transaction{}, &domain.MidtransTransaction{}); err != nil {
		return nil, err
	}

	return db, nil
}
