package usecase

import "speakerseeker-backend/domain"

type RatingUsecase struct {
	ratingRepository domain.RatingRepository
}

func NewRatingUsecase(rr domain.RatingRepository) domain.RatingUsecase {
	return &RatingUsecase{ratingRepository: rr}
}

func (u *RatingUsecase) Create(input domain.CreateRating) error {
	rating := domain.Rating{
		Text:      input.Text,
		Rating:    input.Rating,
		UserId:    input.UserId,
		SpeakerId: input.SpeakerId,
	}
	if err := u.ratingRepository.Create(rating); err != nil {
		return err
	}
	return nil
}
