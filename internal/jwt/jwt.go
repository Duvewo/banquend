package jwt

import (
	"github.com/Duvewo/banquend/models"
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	models.AuthModel
	jwt.RegisteredClaims `json:"claims,omitempty"`
}
