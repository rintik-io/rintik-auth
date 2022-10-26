package role

import (
	"fmt"

	"github.com/rintik-io/rintik-auth/app"
)

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *Role) TableName() string {
	return "p_role"
}

func (p *Role) GenerateTable() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY,
			name TEXT,
			UNIQUE (name)
		)
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	if err := p.InsertDefaultData(); err != nil {
		return err
	}

	return nil
}

func (p *Role) InsertDefaultData() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		INSERT OR IGNORE INTO %s (name) 
			VALUES('admin'), ('user')
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	return nil
}
