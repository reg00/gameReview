package dto

import "github.com/Reg00/gameReview/internal/domain/models"

type Review struct {
	ID          int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	GameID      int    `gorm:"index:idx_gameid,unique" json:"gameid"`
	Description string `gorm:"not null" json:"description"`
	PlayTime    int    `gorm:"not null" json:"playtime"`
	PlayMinutes int    `gorm:"not null" json:"playminutes"`
	Rate        int    `gorm:"not null" json:"rate"`
}

func (review *Review) Convert() *models.GetReview {
	return &models.GetReview{
		ID:          review.ID,
		GameID:      review.GameID,
		Description: review.Description,
		PlayTime:    review.PlayTime,
		PlayMinutes: review.PlayMinutes,
		Rate:        review.Rate,
	}
}

func ConvertToDto(review *models.AddReview) *Review {
	return &Review{
		GameID:      review.GameID,
		Description: review.Description,
		PlayTime:    review.PlayTime,
		PlayMinutes: review.PlayMinutes,
		Rate:        review.Rate,
	}
}
