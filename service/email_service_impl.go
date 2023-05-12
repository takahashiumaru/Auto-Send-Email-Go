package service

import (
	"auto-emails/auth"
	"bytes"
	"encoding/base64"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

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
    // membaca file lampiran
    lampiranPath := "file/ins.png"
	
	// alamat penerima
	toEmail1 := "-"
	toEmail2 := "-"
	to := []string{toEmail1, toEmail2}
    
	// subject
	judul := "Messange Month"
    
	// message
	pesan := "Auto Proses Email"
    
	lampiranBytes, err := os.ReadFile(lampiranPath)
    if err != nil {
        return err
    }

    // data pengirim
    from := "-"
    password := "-"


    // Pengaturan email
	subjek := "Subject: " + judul + "\n"

	// Membuat email buffer
	var email bytes.Buffer
	email.WriteString("From: " + from + "\r\n")
	email.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	email.WriteString(subjek)
	email.WriteString("MIME-version: 1.0\r\n")
	email.WriteString("Content-Type: multipart/mixed; boundary=\"myboundary\"\r\n")
	email.WriteString("\r\n")

	// Menambahkan konten teks
	email.WriteString("--myboundary\r\n")
	email.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	email.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
	email.WriteString("\r\n")
	email.WriteString(pesan)
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
    err = smtp.SendMail(smtpAddr, authEmail, from, to, email.Bytes())
    if err != nil {
        return err
    }
    return nil
}
