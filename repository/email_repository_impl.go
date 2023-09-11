package repository

import (
	"auto-emails/helper"
	"auto-emails/model/domain"

	"gorm.io/gorm"
)

type EmailRepositoryImpl struct {
}

func NewEmailRepository() EmailRepository {
	return &EmailRepositoryImpl{}
}

func (repository *EmailRepositoryImpl) FindOutletMYSQL(tx *gorm.DB, date string) *domain.Emails {
	data := &domain.Emails{}
	err := tx.Table("outlets").Select("id as id,name as name").Where("DATE(created_at) = ?", date).Order("created_at DESC").
		Find(data).Error
	helper.PanicIfError(err)
	return data
}

func (repository *EmailRepositoryImpl) FindOutletMSSQL(tx *gorm.DB, date string) *domain.Emails {
	var data *domain.Emails
	err := tx.Table("M_Outlet").
		Select("CONCAT(MO_KodeArea, MO_KodeOutlet) as ID, MO_NamaOutlet as Name").
		Where("CONVERT(date, MO_CreateTime) = ?", date).
		Order("MO_CreateTime DESC").
		Find(&data).Error

	helper.PanicIfError(err)
	return data
}

func (repository *EmailRepositoryImpl) FindCustomerMYSQL(tx *gorm.DB, date string) *domain.Emails {
	data := &domain.Emails{}
	err := tx.Table("customers").Select("id as id,name as name").Where("DATE(created_at) = ?", date).Order("created_at DESC").
		Find(data).Error
	helper.PanicIfError(err)
	return data
}

func (repository *EmailRepositoryImpl) FindCustomerMSSQL(tx *gorm.DB, customerID []string) *domain.Emails {
	var data *domain.Emails
	err := tx.Table("M_HeaderCustomer").
		Select("MC_KodeCustAll as ID, MC_NamaCust as Name").
		Where("MC_KodeCustAll", customerID).
		Order("MC_UpdateTime DESC").
		Find(&data).Error

	helper.PanicIfError(err)
	return data
}
