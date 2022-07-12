package entity

type Company struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4()" json:"-"`
	Name       string
	LocationID string    `json:"-" gorm:"type:uuid;not null"`
	Location   *Location ` gorm:"foreignkey:LocationID;association_foreignkey:ID"`
	Employees  []Employee
}
