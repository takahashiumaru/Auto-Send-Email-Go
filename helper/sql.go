package helper

import (
	"os"
	"strings"

	"gorm.io/gorm"
)

func RunSQLFromFile(db *gorm.DB, filePath string) {
	file, err := os.ReadFile(filePath)
	PanicIfError(err)

	requests := strings.Split(string(file), "--")
	for _, request := range requests {
		if request != "" {
			err = db.Exec(request).Error
			PanicIfError(err)
		}
	}
}
