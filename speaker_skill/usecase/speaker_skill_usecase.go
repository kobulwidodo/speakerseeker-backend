package usecase

import "speakerseeker-backend/domain"

type SpeakerSkillUsecase struct {
	speakerSkillRepository domain.SpeakerSkillRepository
}

func NewSpeakerSkillUsecase(ssr domain.SpeakerSkillRepository) domain.SpeakerSkillUsecase {
	return &SpeakerSkillUsecase{speakerSkillRepository: ssr}
}

func (u *SpeakerSkillUsecase) GetBySpeakerId(id uint) ([]domain.SpeakerSkill, error) {
	skills, err := u.speakerSkillRepository.FindBySpeakerId(id)
	if err != nil {
		return skills, err
	}
	return skills, nil
}
