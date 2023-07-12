package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Reg00/gameReview/internal/domain/models"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	rdb *redis.Client
	ctx context.Context
}

func Register(cfg *config.Configuration) (*RedisStorage, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Addr,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisStorage{
		rdb: rdb,
		ctx: context.Background(),
	}, nil
}

func (storage *RedisStorage) SetGame(game models.Game) error {
	json, err := json.Marshal(game)
	if err != nil {
		return err
	}

	cacheErr := storage.rdb.Set(storage.ctx, strconv.Itoa(game.ID), json, 24*time.Hour).Err()
	if cacheErr != nil {
		return cacheErr
	}

	return nil
}

func (storage *RedisStorage) GetGameById(id int) (models.Game, error) {
	val, err := storage.rdb.Get(storage.ctx, strconv.Itoa(id)).Bytes()
	if err != nil {
		return models.Game{}, err
	}

	game := toJson(val)

	return game, nil
}

func toJson(val []byte) models.Game {
	game := models.Game{}
	err := json.Unmarshal(val, &game)
	if err != nil {
		panic(err)
	}
	return game
}
