package role

import (
	"fmt"
	"strings"

	"github.com/rintik-io/rintik-auth/app"
)

func (p *Role) GetByFilter(filter []string, operator string) ([]Role, int, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	if operator == "" || operator == " " {
		return nil, 0, fmt.Errorf("operator must not empty")
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, p.TableName(), strings.Join(filter, operator))
	row, err := sqliteConn.Query(query)
	if err != nil {
		return nil, 0, err
	}
	defer row.Close()

	var listRoles []Role
	for row.Next() {
		var currentData Role
		if err = row.Scan(&currentData.ID, &currentData.Name); err != nil {
			return nil, 0, err
		}

		listRoles = append(listRoles, currentData)
	}

	return listRoles, len(listRoles), nil
}

func (p *Role) GetOneByFilter(filter []string, operator string) (Role, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	if operator == "" || operator == " " {
		return Role{}, fmt.Errorf("operator must not empty")
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, p.TableName(), strings.Join(filter, operator))
	row, err := sqliteConn.Query(query)
	if err != nil {
		return Role{}, err
	}
	defer row.Close()

	var output Role
	for row.Next() {
		var currentData Role
		if err = row.Scan(&currentData.ID, &currentData.Name); err != nil {
			return Role{}, err
		}

		output = currentData
	}

	return output, nil
}
