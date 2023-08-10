package jwt

import (
	"fmt"
	"time"

	gJWT "github.com/golang-jwt/jwt/v5"
)

const (
	DATE_FORMAT = "02.01.2006 15:04:05"
)

func New(jData map[string]interface{}, secret []byte, longTime time.Duration) (string, error) {

	jwtData := gJWT.MapClaims{
		"expires_in": time.Now().UTC().Add(time.Minute * longTime).Format(DATE_FORMAT),
	}

	for k, v := range jData {
		jwtData[k] = v
	}

	token := gJWT.NewWithClaims(gJWT.SigningMethodHS256, jwtData)
	return token.SignedString(secret)
}

func IsExpired(token *gJWT.Token) (bool, error) {
	if claims, ok := token.Claims.(gJWT.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		if _, ok := claims["expires_in"]; !ok {
			return true, fmt.Errorf("Token not valid")
		}

		expired_dt, err := time.Parse(DATE_FORMAT, claims["expires_in"].(string))
		if err != nil {
			return true, fmt.Errorf("Not found attribute expires_in")
		}
		if expired_dt.Sub(time.Now().UTC()) < time.Duration(0) {
			return true, nil
		}
	}

	return false, nil

}

func Check(hToken string, secret []byte) (*gJWT.Token, error) {

	token, err := gJWT.Parse(hToken, func(token *gJWT.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*gJWT.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
