package users

import (
	"fmt"

	"github.com/rintik-io/rintik-auth/app"
)

func (p *Users) Create() error {
	sqliteConn := app.Properties.Databases.SQLiteConn
	sqlStatement := fmt.Sprintf(`
		INSERT INTO %s 
			(email, username, phone, name, role_id, password, created_at) 
		VALUES
			($1, $2, $3, $4, $5, $6, $7)`, p.TableName())
	_, err := sqliteConn.Exec(sqlStatement, p.Email, p.Username, p.Phone, p.Name, p.RoleID, p.Password, p.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
