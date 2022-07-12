package entity

type Location struct {
	ID        string `json:"-" gorm:"type:uuid;default:uuid_generate_v4()"`
	Longitude int
	Latitude  int
	City      string
	Companies []Company `gorm:"foreignKey:LocationID"`
}
