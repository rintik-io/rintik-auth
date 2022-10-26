package users

import (
	"fmt"

	"github.com/rintik-io/rintik-auth/app"
)

type RequestRegister struct {
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Role     string `form:"role" binding:"required"`
}

type RequestClaims struct {
	UsernameEmailPhone string `form:"username_email_phone" binding:"required"`
	Password           string `form:"password" binding:"required"`
}

type RequestValidate struct {
	Jwt string `form:"jwt" binding:"required"`
}

type Users struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Password  string `json:"password,omitempty"`
	RoleID    int    `json:"role_id"`
	CreatedAt string `json:"created_at"`
}

func (p *Users) TableName() string {
	return "t_users"
}

func (p *Users) KeyRedis() string {
	return "data:users"
}

func (p *Users) GenerateTable() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY,
			email TEXT,
			username TEXT,
			phone TEXT,
			name TEXT,
			password TEXT,
			role_id INTEGER,
			created_at TEXT
		)
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	return nil
}
