package core

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type TokenType byte


const (
	USER_TOKEN_TYPE_NORMAL                         TokenType = 1
	USER_TOKEN_TYPE_REQUIRE_2FA                    TokenType = 2
	USER_TOKEN_TYPE_EMAIL_VERIFY                   TokenType = 3
	USER_TOKEN_TYPE_PASSWORD_RESET                 TokenType = 4
	USER_TOKEN_TYPE_MCP                            TokenType = 5
	USER_TOKEN_TYPE_OAUTH2_CALLBACK_REQUIRE_VERIFY TokenType = 6
	USER_TOKEN_TYPE_OAUTH2_CALLBACK                TokenType = 7
	USER_TOKEN_TYPE_API                            TokenType = 8
)


type UserTokenClaims struct {
	UserTokenId string    `json:"userTokenId"`
	Uid         int64     `json:"jti,string"`
	Username    string    `json:"username,omitempty"`
	Type        TokenType `json:"type"`
	IssuedAt    int64     `json:"iat"`
	ExpiresAt   int64     `json:"exp"`
}


func (c *UserTokenClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: time.Unix(c.ExpiresAt, 0),
	}, nil
}


func (c *UserTokenClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: time.Unix(c.IssuedAt, 0),
	}, nil
}


func (c *UserTokenClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{}, nil
}


func (c *UserTokenClaims) GetIssuer() (string, error) {
	return "", nil
}


func (c *UserTokenClaims) GetSubject() (string, error) {
	return "", nil
}


func (c *UserTokenClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
