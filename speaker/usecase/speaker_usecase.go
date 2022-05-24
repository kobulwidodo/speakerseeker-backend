package usecase

import "speakerseeker-backend/domain"

type SpeakerUsecase struct {
	speakerRepository      domain.SpeakerRepository
	speakerSkillRepository domain.SpeakerSkillRepository
}

func NewSpeakserUsecase(sr domain.SpeakerRepository, ssr domain.SpeakerSkillRepository) domain.SpeakerUsecase {
	return &SpeakerUsecase{speakerRepository: sr, speakerSkillRepository: ssr}
}

func (u *SpeakerUsecase) GetAll(query string) ([]domain.Speaker, error) {
	speakers, err := u.speakerRepository.FindAll(query)
	if err != nil {
		return speakers, err
	}
	return speakers, nil
}

func (u *SpeakerUsecase) GetById(id uint) (domain.Speaker, error) {
	speaker, err := u.speakerRepository.FindOne(id)
	if err != nil {
		return speaker, err
	}
	return speaker, nil
}
