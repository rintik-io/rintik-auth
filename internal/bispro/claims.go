package bispro

import (
	"fmt"
	"time"

	"github.com/fahmyabdul/golibs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rintik-io/rintik-auth/configs"
	"github.com/rintik-io/rintik-auth/internal/models/users"
	"golang.org/x/crypto/bcrypt"
)

type BisproClaims struct{}

func (p *BisproClaims) JwtConf() map[string]interface{} {
	conf, ok := configs.Properties.Etc["jwt"].(map[string]interface{})
	if !ok {
		return nil
	}

	return conf
}

func (p *BisproClaims) Claims(requestData *users.RequestClaims) (gin.H, error) {
	var userModel users.Users

	userData, err := userModel.GetOneByFilter([]string{
		fmt.Sprintf("username = '%s'", requestData.UsernameEmailPhone),
		fmt.Sprintf("email = '%s'", requestData.UsernameEmailPhone),
		fmt.Sprintf("phone = '%s'", requestData.UsernameEmailPhone),
	}, "OR")
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | GetUserByFilter | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, err
	}

	if userData.Phone == "" {
		golibs.Log.Printf("| Bispro | Register | GetUserByFilter | Failed, error: Unauthorized | Request: %v", requestData)
		return nil, fmt.Errorf("Unauthorized")
	}

	// Compare password with passwordHashed
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(requestData.Password))
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | CompareHashAndPassword | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, fmt.Errorf("Unauthorized")
	}

	expSeconds := p.JwtConf()["exp"].(int)
	tokenExpDate := time.Now().Add(time.Second * time.Duration(expSeconds)).Unix()

	userData.Password = ""
	userData.ID = 0

	claims := &jwt.MapClaims{
		"exp":     tokenExpDate,
		"content": userData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(p.JwtConf()["secret"].(string))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | CreateToken | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, fmt.Errorf("Unauthorized")
	}

	return gin.H{"jwt": tokenString}, nil
}
