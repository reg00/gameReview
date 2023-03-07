package igdb

import (
	"fmt"
	"log"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/Reg00/gameReview/internal/domain/dto"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
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

func (client *IgdbClient) GetGamesByName(offset int, limit int, name string) ([]dto.Game, error) {
	var gms []dto.Game

	games, err := client.Client.Games.Search(
		name,
		igdb.SetFields("name", "cover", "genres"),
		igdb.SetOffset(offset),
		igdb.SetLimit(limit),
		//igdb.SetOrder("name", igdb.OrderAscending),
	)

	if err != nil {
		fmt.Println("Error while searching games: " + err.Error())

		return nil, err
	}

	for _, game := range games {
		var img string
		var genrs []string

		if game.Cover != 0 {
			cover, err := client.Client.Covers.Get(game.Cover, igdb.SetFields("image_id"))
			if err != nil {
				log.Println(err)
			}

			img, err = cover.SizedURL(igdb.SizeCoverSmall, 1)
			if err != nil {
				log.Println(err)
			}
		}

		if len(game.Genres) > 0 {
			genres, err := client.Client.Genres.List(game.Genres, igdb.SetFields("name"))
			if err != nil {
				log.Println(err)
			}

			for _, genre := range genres {
				genrs = append(genrs, genre.Name)
			}
		}

		gms = append(gms, dto.Game{
			Name:     game.Name,
			ImageURI: img,
			Genres:   genrs,
		})
	}

	return gms, nil
}
