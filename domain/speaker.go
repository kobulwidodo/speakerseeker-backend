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
	Linkedin           string              `json:"linkedid"`
	ImgPath            string              `json:"img_path"`
	SpeakerSkills      []SpeakerSkill      `json:"speaker_skills"`
	SpeakerCareers     []SpeakerCareer     `json:"speaker_careers"`
	SpeakerExperiences []SpeakerExperience `json:"speaker_experiences"`
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
