package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"auto-emails/exception"
	"auto-emails/helper"

	c "auto-emails/configuration"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CreateAuthFunc func(userID uint, tokenDetails *TokenDetails)

type AccessDetails struct {
	UserID   uint
	Role     string
	Level    string
	Nip      string
	Name     string
	MainID   uint
	MainRole string
}

type NewAccessDetails struct {
	Nip  string
	Name string
	Role string
	ID   uint
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	UserAgent           string
	RemoteAddress       string
	AtExpired           int64
	RefreshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(ExtractToken(c.Request))
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Check Permission User
		// if !helper.Contains(roles, tokenAuth.Role) {
		// 	helper.PanicIfError(exception.ErrPermissionDenied)
		// }

		next(c, tokenAuth)
	}
}

func ExtractTokenMetadata(stringToken string) (*AccessDetails, error) {
	token, err := VerifyToken(stringToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		mainId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["main_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			UserID:   uint(userId),
			Role:     claims["role"].(string),
			Level:    claims["level"].(string),
			Nip:      claims["nip"].(string),
			Name:     claims["name"].(string),
			MainRole: claims["main_role"].(string),
			MainID:   uint(mainId),
		}, nil
	}
	return nil, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	configuration, err := c.LoadConfig()
	if err != nil {
		return nil, err
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configuration.AccessSecret), nil
	})
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
