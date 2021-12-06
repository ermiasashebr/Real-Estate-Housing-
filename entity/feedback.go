package entity

type Feedback struct {
	ID	uint
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"type:varchar(100);not null"`
	Description string
	Image string `gorm:"type:varchar(255)"`
	UserId uint `json:"userId" gorm:"not null"`
}