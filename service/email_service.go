package service

import (
	"auto-emails/auth"
)

type EmailService interface {
	EmailProsess(auth *auth.AccessDetails) error
}
