package igdb

import (
	"fmt"
	"strings"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/Reg00/gameReview/internal/domain/dto/httperr"
	"github.com/Reg00/gameReview/internal/domain/models"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/pkg/errors"
)

type IgdbClient struct {
	Client *igdb.Client
}

func Register(
	cfg *config.Configuration,
) *IgdbClient {
	igdbClient := new(IgdbClient)
	igdbClient.Client = igdb.NewClient(cfg.Igdb.ClientId, cfg.Igdb.ClientSecret, nil)

	return igdbClient
}

func (client *IgdbClient) GetGameById(id int) (models.Game, error) {
	game, err := client.Client.Games.Get(id, igdb.SetFields("name", "cover", "genres"))
	if err != nil {
		return models.Game{}, handleError(err)
	}

	dtoGame, err := client.convertToDto(game)
	if err != nil {
		return models.Game{}, handleError(err)
	}

	return dtoGame, nil
}

func (client *IgdbClient) GetGamesByName(offset int, limit int, name string) ([]models.Game, error) {
	var gms []models.Game

	opts := igdb.ComposeOptions(
		igdb.SetFields("name", "cover", "genres"),
		igdb.SetOffset(offset),
		igdb.SetLimit(limit))

	var games []*igdb.Game
	var err error

	if name != "" {
		games, err = client.Client.Games.Search(
			name, opts,
		)
	} else {
		games, err = client.Client.Games.Index(
			opts)
	}

	client.Client.Games.Index()

	if err != nil {
		return nil, handleError(err)
	}

	for _, game := range games {

		dtoGame, err := client.convertToDto(game)
		if err != nil {
			return nil, handleError(err)
		}

		gms = append(gms, dtoGame)
	}

	return gms, nil
}

func (client *IgdbClient) convertToDto(game *igdb.Game) (models.Game, error) {
	var img string
	var genrs []string

	if game.Cover != 0 {
		cover, err := client.Client.Covers.Get(game.Cover, igdb.SetFields("image_id"))
		if err != nil {
			return models.Game{}, handleError(err)
		}

		img, err = cover.SizedURL(igdb.SizeCoverSmall, 1)
		if err != nil {
			return models.Game{}, handleError(err)
		}
	}

	if len(game.Genres) > 0 {
		genres, err := client.Client.Genres.List(game.Genres, igdb.SetFields("name"))
		if err != nil {
			return models.Game{}, handleError(err)
		}

		for _, genre := range genres {
			genrs = append(genrs, genre.Name)
		}
	}

	dtoGame := models.Game{
		ID:       game.ID,
		Name:     game.Name,
		ImageURI: img,
		Genres:   genrs,
	}

	return dtoGame, nil
}

func handleError(err error) error {
	if strings.Contains(err.Error(), "cannot get Game with ID") {
		err = errors.Wrap(httperr.ErrNotFound, fmt.Sprintf("IGDB: %s", err.Error()))
	}

	return err
}
