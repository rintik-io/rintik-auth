package bispro

import (
	"fmt"

	"github.com/fahmyabdul/golibs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rintik-io/rintik-auth/configs"
)

type BisproValidate struct{}

func (p *BisproValidate) JwtConf() map[string]interface{} {
	conf, ok := configs.Properties.Etc["jwt"].(map[string]interface{})
	if !ok {
		return nil
	}

	return conf
}

func (p *BisproValidate) Validate(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		secretKey := []byte(p.JwtConf()["secret"].(string))

		return secretKey, nil
	})
	if err != nil {
		golibs.Log.Printf("| Bispro | Validate | Failed, error: %s | Request: %v", err.Error(), jwtToken)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		golibs.Log.Printf("| Bispro | Validate | Failed, error: Invalid Token | Request: %v", jwtToken)
		return nil, err
	}

	return claims, nil
}
