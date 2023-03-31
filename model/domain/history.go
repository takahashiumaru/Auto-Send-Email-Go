package domain

import (

	"gorm.io/gorm"
)

type Histories []History

type History struct {
	// Required Fields
	gorm.Model
	CreatedByID uint `gorm:""`

	// Fields
	TableName string    `gorm:"size:100"`
	TableID   string    `gorm:"size:50"`
	Data      string    `gorm:"size:8000"`
	Type      string    `gorm:"size:10"`
}
