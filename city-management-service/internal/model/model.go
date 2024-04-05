package model

type City struct {
	UUID      string  `bson:"_id,omitempty"`
	Name      string  `gorm:"index:idx_name_country,unique"`
	Country   string  `gorm:"index:idx_name_country,unique"`
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}
