package domain

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Text      string `json:"text"`
	Rating    int    `json:"rating"`
	UserId    uint
	SpeakerId uint
}

type RatingRepository interface {
	Create(rating Rating) error
}

type RatingUsecase interface {
	Create(input CreateRating) error
}

type CreateRating struct {
	Text      string `binding:"required"`
	Rating    int    `binding:"required"`
	UserId    uint
	SpeakerId uint
}

type BindingUriSpeakerId struct {
	Id uint `uri:"id" binding:"required"`
}
