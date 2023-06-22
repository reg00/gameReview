package storage

type Review struct {
	ID          int32  `json:"id"`
	GameID      int32  `json:"gameid"`
	Description string `json:"description"`
	PlayTime    int32  `json:"playtime"`
	Rate        int8   `json:"rate"`
}
