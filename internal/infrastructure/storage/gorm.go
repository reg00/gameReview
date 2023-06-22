package gorm

import (
	"errors"
	"fmt"
	"log"

	"github.com/Reg00/gameReview/internal/domain/dto/storage"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStorage struct {
	gormDB *gorm.DB
}

func Register(cfg *config.Configuration) (*GormStorage, error) {
	opts, err := newGormOptions(cfg.Storage.Options)
	if err != nil {
		return nil, fmt.Errorf("gorm: malformed options: %s", err)
	}
	log.Println(opts.dsn)
	var db = &gorm.DB{}
	switch {
	case opts.driver == "postgres":
		db, err = gorm.Open(postgres.Open(opts.dsn))
		if err != nil {
			return nil, fmt.Errorf("gorm: open db: %s", err)
		}
	default:
		return nil, errors.New("gorm: invalid storage driver")
	}
	db.AutoMigrate(&storage.Review{})
	return &GormStorage{
		gormDB: db,
	}, nil
}

func (storage *GormStorage) AddReview(review *storage.Review) (*storage.Review, error) {
	result := storage.gormDB.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}
