package jwt_generator

import "github.com/golang-jwt/jwt/v4"

const IssuerDefault = "lisansly"

const (
	RoleUser = "user"
)

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
