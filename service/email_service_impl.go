package service

import (
	"auto-emails/auth"
	"auto-emails/helper"
	"fmt"
	"net/smtp"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EmailServiceImpl struct {
	DB *gorm.DB
}

func NewEmailService(
	db *gorm.DB,
	validate *validator.Validate,
) EmailService {
	return &EmailServiceImpl{
		DB: db,
	}
}

func (service *EmailServiceImpl) EmailProsess(auth *auth.AccessDetails) error {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// data pengirim
	from := "---"
	password := "---"

	// alamat penerima
	toEmail1 := "---"
	toEmail2 := "---"
	to := []string{toEmail1, toEmail2}

	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// pesan
	subject := "Subject: Email Pertama Kami\n"
	body := "<h1>Ini adalah isi email pertama kami menggunakan Golang</h1>"

	message := []byte(subject + "\r\n" + body)

	// data otentikasi
	authEmail := smtp.PlainAuth("", from, password, host)

	// kirim email
	err = smtp.SendMail(address, authEmail, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
	}

	return nil
}
