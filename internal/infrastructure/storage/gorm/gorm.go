package gorm

import (
	"fmt"

	dto "github.com/Reg00/gameReview/internal/domain/dto"
	"github.com/Reg00/gameReview/internal/domain/dto/httperr"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStorage struct {
	gormDB *gorm.DB
}

func Register(cfg *config.Configuration) (*GormStorage, error) {
	var db = &gorm.DB{}

	err := createDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("gorm: initial create: %s", err)
	}

	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Storage.Options.Host,
		cfg.Storage.Options.User,
		cfg.Storage.Options.Password,
		cfg.Storage.Options.Dbname,
		cfg.Storage.Options.Port)

	switch {
	case cfg.Storage.Options.Driver == "postgres":
		db, err = gorm.Open(postgres.Open(createDBDsn))
		if err != nil {
			return nil, fmt.Errorf("gorm: open db: %s", err)
		}
	default:
		return nil, errors.New("gorm: invalid storage driver")
	}

	db.AutoMigrate(&dto.Review{})
	return &GormStorage{
		gormDB: db,
	}, nil
}

func createDatabase(cfg *config.Configuration) error {
	conn_url := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		cfg.Storage.Options.User,
		cfg.Storage.Options.Password,
		cfg.Storage.Options.Host,
		cfg.Storage.Options.Port)
	DB, err := gorm.Open(postgres.Open(conn_url), &gorm.Config{})
	if err != nil {
		return err
	}

	count := 0
	DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", cfg.Storage.Options.Dbname).Scan(&count)
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", cfg.Storage.Options.Dbname)
		DB.Exec(sql)
	}

	return nil
}

func (storage *GormStorage) AddReview(review *dto.Review) (*dto.Review, error) {
	result := storage.gormDB.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}

func (storage *GormStorage) UpdateReview(id int, updatReview *dto.Review) (*dto.Review, error) {
	review, err := storage.GetReviewById(id)
	if err != nil {
		return nil, err
	}

	storage.gormDB.Model(&review).Updates(dto.Review{
		Description: updatReview.Description,
		PlayTime:    updatReview.PlayTime,
		PlayMinutes: updatReview.PlayMinutes,
		Rate:        updatReview.Rate,
	})

	review, err = storage.GetReviewById(id)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (storage *GormStorage) DeleteReview(id int) error {
	review, err := storage.GetReviewById(id)
	if err != nil {
		return err
	}

	storage.gormDB.Delete(&review)
	return nil
}

func (storage *GormStorage) GetReviewById(id int) (*dto.Review, error) {
	var review dto.Review
	storage.gormDB.First(&review, "id = ?", id)

	if review == (dto.Review{}) {
		return &dto.Review{}, errors.Wrap(httperr.ErrNotFound, fmt.Sprintf("GORM: %s", errors.New(fmt.Sprintf("Cannot find review with %d", id))))
	}

	return &review, nil
}
