package storage

type Review struct {
	ID          int32  `gorm:"primaryKey;autoIncrement:true" json:"id"`
	GameID      int32  `gorm:"index:idx_gameid,unique" json:"gameid"`
	Description string `gorm:"not null" json:"description"`
	PlayTime    int32  `gorm:"not null" json:"playtime"`
	PlayMinutes int32  `gorm:"not null" json:"playminutes"`
	Rate        int8   `gorm:"not null" json:"rate"`
}
