package repository

import (
	"auto-emails/model/domain"

	"gorm.io/gorm"
)

type EmailRepository interface {
	FindOutletMYSQL(tx *gorm.DB, date string) *domain.Emails
	FindOutletMSSQL(tx *gorm.DB, date string) *domain.Emails
	FindCustomerMYSQL(tx *gorm.DB, date string) *domain.Emails
	FindCustomerMSSQL(tx *gorm.DB, customerID []string) *domain.Emails
}
