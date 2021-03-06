package domain

import "gorm.io/gorm"

type Speaker struct {
	gorm.Model
	Name               string              `json:"name"`
	HeaderTitle        string              `json:"header_title"`
	Header             string              `json:"header"`
	About              string              `gorm:"text" json:"about"`
	Location           string              `json:"location"`
	VirtualFeeStart    int32               `json:"virtual_fee_start"`
	VirtualFeeStop     int32               `json:"virtual_fee_stop"`
	SiteFeeStart       int32               `json:"site_fee_start"`
	SiteFeeStop        int32               `json:"site_fee_stop"`
	Tiktok             string              `json:"tiktok"`
	Instagram          string              `json:"instagram"`
	Linkedin           string              `json:"linkedin"`
	ImgPath            string              `json:"img_path"`
	ImgProfilePath     string              `json:"img_profile_path"`
	ImgReviewPath      string              `json:"img_review_path"`
	SpeakerSkills      []SpeakerSkill      `json:"speaker_skills"`
	SpeakerCareers     []SpeakerCareer     `json:"speaker_careers"`
	SpeakerExperiences []SpeakerExperience `json:"speaker_experiences"`
	Transactions       []Transaction       `json:"transaction"`
}

type SpeakerRepository interface {
	FindAll(query string) ([]Speaker, error)
	FindOne(id uint) (Speaker, error)
}

type SpeakerUsecase interface {
	GetAll(query string) ([]Speaker, error)
	GetById(id uint) (Speaker, error)
}

type UriById struct {
	Id uint `uri:"id" binding:"required"`
}
