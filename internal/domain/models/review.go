package models

type GetReview struct {
	ID          int
	Description string
	PlayTime    int
	PlayMinutes int
	Rate        int
	Game        Game
}

type AddReview struct {
	GameID      int
	Description string
	PlayTime    int
	PlayMinutes int
	Rate        int
}