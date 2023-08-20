package helpers

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Exp   string `json:"exp"`
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, expired int) (string, error) {
	duration, _ := strconv.Atoi(LoadEnv("JWT_TIME_DURATION"))
	if expired > 0 {
		duration = expired
	}

	drt := time.Minute * time.Duration(duration)
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(drt).Unix()

	for i, v := range Data {
		claims[i] = v
	}
	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(LoadEnv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func VerifyTokenHeader(requestToken string) (MetaToken, error) {

	token, err := jwt.Parse((requestToken), func(token *jwt.Token) (interface{}, error) {
		return []byte(LoadEnv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println(err)
		return MetaToken{}, err
	}
	claimToken := DecodeToken(token)
	return claimToken.Claims, nil
}

func VerifyToken(accessToken, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := LoadEnv(SecrePublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
