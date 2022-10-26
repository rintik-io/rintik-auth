package bispro

import (
	"fmt"
	"time"

	"github.com/fahmyabdul/golibs"
	"github.com/rintik-io/rintik-auth/internal/common"
	"github.com/rintik-io/rintik-auth/internal/models/role"
	"github.com/rintik-io/rintik-auth/internal/models/users"
)

type BisproRegister struct{}

func (p *BisproRegister) Register(requestData *users.RequestRegister) (*users.Users, error) {
	var (
		userModel users.Users
		roleModel role.Role
	)

	_, countUsers, err := userModel.GetByFilter([]string{
		fmt.Sprintf("phone = '%s'", requestData.Phone),
		fmt.Sprintf("email = '%s'", requestData.Email),
		fmt.Sprintf("username = '%s'", requestData.Username),
	}, "OR")
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | GetUserByFilter | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, err
	}

	if countUsers > 0 {
		golibs.Log.Printf("| Bispro | Register | GetUSerByFilter | Failed, error: phone/email/username already exists | Request: %v", requestData)
		return nil, fmt.Errorf("phone/email/username already exists")
	}

	role, err := roleModel.GetOneByFilter([]string{
		fmt.Sprintf("name = '%s'", requestData.Role),
	}, "AND")
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | GetRoleByFilter | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, err
	}
	if role.Name == "" {
		golibs.Log.Printf("| Bispro | Register | GetRoleByFilter | Failed, error: role '%s' not found | Request: %v", requestData.Role, requestData)
		return nil, fmt.Errorf("role '%s' not found", requestData.Role)
	}

	// var genPassword string
	// passwordCheckStatus := true
	// for passwordCheckStatus == true {
	// 	genPassword = common.GenerateRandPassword(10)
	// 	_, countPassword, err := userModel.GetByFilter([]string{
	// 		fmt.Sprintf("password = '%s'", genPassword),
	// 	})
	// 	if err != nil {
	// 		golibs.Log.Printf("| Bispro | Register | GenerateRandPassword | Failed, error: %s | Request: %v", err.Error(), requestData)
	// 		return nil, err
	// 	}

	// 	if countPassword == 0 {
	// 		passwordCheckStatus = false
	// 	} else {
	// 		time.Sleep(time.Duration(1) * time.Second)
	// 	}
	// }

	passwordHashed, err := common.GenerateBcryptPassword(requestData.Password, 16)
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | GenPassword | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, err
	}

	newUser := users.Users{
		Email:     requestData.Email,
		Phone:     requestData.Phone,
		Username:  requestData.Username,
		Name:      requestData.Name,
		RoleID:    role.ID,
		Password:  passwordHashed,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = newUser.Create()
	if err != nil {
		golibs.Log.Printf("| Bispro | Register | CreateUser | Failed, error: %s | Request: %v", err.Error(), requestData)
		return nil, err
	}

	golibs.Log.Printf("| Bispro | Register | Success | Request: %v", requestData)

	newUser.Password = ""

	return &newUser, nil
}
