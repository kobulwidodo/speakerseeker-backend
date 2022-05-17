package domain

import "gorm.io/gorm"

type Speaker struct {
	gorm.Model
	Name            string
	HeaderTitle     string
	Header          string
	About           string `gorm:"text"`
	VirtualFeeStart int32
	VirtualFeeStop  int32
	SiteFeeStart    int32
	SiteFeeStop     int32
	Tiktok          string
	Instagram       string
	Linkedin        string
	ImgPath         string
}

type SpekaerRepository interface {
	FindAll() ([]Speaker, error)
	FindOne(id uint) (Speaker, error)
}

type SpeakerUsecase interface {
	GetAll() ([]Speaker, error)
	GetById(id uint) (Speaker, error)
}

type UriById struct {
	Id uint `uri:"id" binding:"required"`
}
