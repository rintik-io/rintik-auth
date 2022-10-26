package initialization

import (
	"github.com/rintik-io/rintik-auth/internal/models/role"
	"github.com/rintik-io/rintik-auth/internal/models/users"
)

func InitBefore() error {
	// Generate Table Users
	modelUsers := users.Users{}
	err := modelUsers.GenerateTable()
	if err != nil {
		return err
	}
	// Generate Table Role
	modelRole := role.Role{}
	err = modelRole.GenerateTable()
	if err != nil {
		return err
	}

	return nil
}

func InitAfter() error {
	return nil
}
