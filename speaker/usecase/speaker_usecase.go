package usecase

import "speakerseeker-backend/domain"

type SpeakerUsecase struct {
	speakerRepository domain.SpekaerRepository
}

func NewSpeakserUsecase(sr domain.SpekaerRepository) domain.SpeakerUsecase {
	return &SpeakerUsecase{speakerRepository: sr}
}

func (u *SpeakerUsecase) GetAll() ([]domain.Speaker, error) {
	speakers, err := u.speakerRepository.FindAll()
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
