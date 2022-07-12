package entity

type Employee struct {
	ID         string `json:"-" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string
	Age        int
	Gender     string
	CompanyID  string
	LocationID string    `json:"-" gorm:"type:uuid;not null"`
	Location   *Location ` gorm:"foreignkey:LocationID;association_foreignkey:ID"`
}
