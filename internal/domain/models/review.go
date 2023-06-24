package models

type GetReview struct {
	ID          int
	GameID      int
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

type UpdateReview struct {
	Description string
	PlayTime    int
	PlayMinutes int
	Rate        int
}
