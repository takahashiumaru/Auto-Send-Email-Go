package domain

type Emails []Email
type Email struct {
	// Required Fields
	ID   string `gorm:"column:id;"`
	Name string `gorm:"column:name;"`
}
