package service

import (
	"auto-emails/auth"
	"auto-emails/helper"
	"auto-emails/model/domain"
	"auto-emails/repository"
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	c "auto-emails/configuration"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EmailServiceImpl struct {
	EmailRepository repository.EmailRepository
	DBMS            *gorm.DB
	DBMY            *gorm.DB
	Validate        *validator.Validate
}

func NewEmailService(
	emailRepository repository.EmailRepository,
	dbMS *gorm.DB,
	dbMY *gorm.DB,
	validate *validator.Validate,
) EmailService {
	return &EmailServiceImpl{
		EmailRepository: emailRepository,
		DBMS:            dbMS,
		DBMY:            dbMY,
		Validate:        validate,
	}
}

func (service *EmailServiceImpl) EmailProsess(auth *auth.AccessDetails) error {
	var err error
	txMs := service.DBMS.Begin()
	err = txMs.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(txMs)

	txMy := service.DBMY.Begin()
	err = txMy.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(txMy)

	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	from := configuration.FromEmail
	password := configuration.PasswordEmail

	to := []string{"umar.muataqin@vneu.co.id"}
	lampiranPath := "file/go.jpeg"

	timeNow := time.Now().Format("2006-01-02")
	formattedTime := time.Now().Format("2006-01-02 15:04:05")

	outletMssqlArray := []string{}
	outletMysqlArray := []string{}
	outletMysql := service.EmailRepository.FindOutletMYSQL(txMy, timeNow)
	outletMssql := service.EmailRepository.FindOutletMSSQL(txMs, timeNow)

	// Populate outletMssqlArray
	for _, data := range *outletMssql {
		outletMssqlArray = append(outletMssqlArray, data.ID)
	}

	// Populate outletMysqlArray
	for _, data := range *outletMysql {
		outletMysqlArray = append(outletMysqlArray, data.ID)
	}

	oldArray := []string{}
	arrayTwoMap := make(map[string]bool)
	for _, item := range outletMssqlArray {
		oldArray = append(oldArray, item)
		arrayTwoMap[item] = true
	}

	// Create a new array to store elements from arrayOne that are not in arrayTwo
	newArray := []string{}
	for _, item := range outletMysqlArray {
		if !arrayTwoMap[item] {
			newArray = append(newArray, item)
		}
	}

	// subject
	judul := fmt.Sprintf(`Report Process Master Outlet Tanggal %s`, formattedTime)

	lampiranBytes, err := os.ReadFile(lampiranPath)
	if err != nil {
		return err
	}

	// Pengaturan email
	subject := "Subject: " + judul + "\n"

	// message
	message := ""
	if len(newArray) > 0 {
		message += "Outlet dengan kode ini tidak berhasil:\n"
		message += strings.Join(newArray, "\n")
	} else {
		message += "Semua Outlet berhasil ke insert:\n"
		message += strings.Join(oldArray, "\n")
	}
	err = processEmail(from, to, subject, message, lampiranPath, lampiranBytes, password)
	helper.PanicIfError(err)
	customerMysqlArray := []string{}
	customerMssqlArray := []string{}

	// Fetch customer data from MySQL database
	customerMysql := service.EmailRepository.FindCustomerMYSQL(txMy, timeNow)

	// Populate customerMysqlArray
	for _, data := range *customerMysql {
		customerMysqlArray = append(customerMysqlArray, data.ID)
	}

	var customerMssql *domain.Emails
	if len(customerMysqlArray) > 0 {
		// Fetch customer data from MSSQL database using the IDs from customerMysqlArray
		customerMssql = service.EmailRepository.FindCustomerMSSQL(txMs, customerMysqlArray)
		// Populate customerMssqlArray
		for _, data := range *customerMssql {
			customerMssqlArray = append(customerMssqlArray, data.ID)
		}
	}

	// Create an old customer array and a map for efficient lookup
	oldCustomerArray := []string{}
	customerTwoMap := make(map[string]bool)

	// Populate the oldCustomerArray and customerTwoMap
	for _, item := range customerMssqlArray {
		oldCustomerArray = append(oldCustomerArray, item)
		customerTwoMap[item] = true
	}

	// Create a new customer array to store elements from customerMysqlArray that are not in customerMssqlArray
	newCustomerArray := []string{}
	for _, item := range customerMysqlArray {
		if !customerTwoMap[item] {
			newCustomerArray = append(newCustomerArray, item)
		}
	}

	// Subject
	judul = fmt.Sprintf("Report Process Master Customer Tanggal %s", formattedTime)

	// Read the attachment file
	lampiranBytes, err = os.ReadFile(lampiranPath)
	if err != nil {
		return err
	}

	// Email configuration
	subject = "Subject: " + judul + "\n"

	// Message
	message = ""
	if len(newCustomerArray) > 0 {
		message += "Customer dengan kode ini tidak berhasil:\n"
		message += strings.Join(newCustomerArray, "\n")
	} else {
		message += "Semua Customer berhasil ke insert:\n"
		message += strings.Join(oldCustomerArray, "\n")
	}

	// Send the email
	err = processEmail(from, to, subject, message, lampiranPath, lampiranBytes, password)
	helper.PanicIfError(err)

	return nil
}

func processEmail(from string, to []string, subject string, message string, lampiranPath string, lampiranBytes []byte, password string) error {

	// Membuat email buffer
	var email bytes.Buffer
	email.WriteString("From: " + from + "\r\n")
	email.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	email.WriteString(subject)
	email.WriteString("MIME-version: 1.0\r\n")
	email.WriteString("Content-Type: multipart/mixed; boundary=\"myboundary\"\r\n")
	email.WriteString("\r\n")

	// Menambahkan konten teks
	email.WriteString("--myboundary\r\n")
	email.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	email.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
	email.WriteString("\r\n")
	email.WriteString(message)
	email.WriteString("\r\n")

	// Menambahkan lampiran gambar
	email.WriteString("--myboundary\r\n")
	email.WriteString("Content-Type: image/png; name=\"" + filepath.Base(lampiranPath) + "\"\r\n")
	email.WriteString("Content-Transfer-Encoding: base64\r\n")
	email.WriteString("Content-Disposition: attachment; filename=\"" + filepath.Base(lampiranPath) + "\"\r\n")
	email.WriteString("\r\n")
	encodedLampiran := base64.StdEncoding.EncodeToString(lampiranBytes)
	email.WriteString(encodedLampiran)
	email.WriteString("\r\n--myboundary--\r\n")

	// pengaturan koneksi SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpAddr := smtpHost + ":" + smtpPort
	authEmail := smtp.PlainAuth("", from, password, smtpHost)

	// mengirim email
	err := smtp.SendMail(smtpAddr, authEmail, from, to, email.Bytes())
	if err != nil {
		return err
	}
	return nil
}
