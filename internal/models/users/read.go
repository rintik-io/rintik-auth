package users

import (
	"fmt"
	"strings"

	"github.com/rintik-io/rintik-auth/app"
)

func (p *Users) GetByFilter(filter []string, operator string) ([]Users, int, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	if operator == "" || operator == " " {
		return nil, 0, fmt.Errorf("operator must not empty")
	}

	operator = fmt.Sprintf(" %s ", operator)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, p.TableName(), strings.Join(filter, operator))
	row, err := sqliteConn.Query(query)
	if err != nil {
		return nil, 0, err
	}
	defer row.Close()

	var output []Users
	for row.Next() {
		var currentData Users
		if err = row.Scan(
			&currentData.ID,
			&currentData.Email,
			&currentData.Username,
			&currentData.Phone,
			&currentData.Name,
			&currentData.Password,
			&currentData.RoleID,
			&currentData.CreatedAt,
		); err != nil {
			return nil, 0, err
		}

		output = append(output, currentData)
	}

	return output, len(output), nil
}

func (p *Users) GetOneByFilter(filter []string, operator string) (Users, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	if operator == "" || operator == " " {
		return Users{}, fmt.Errorf("operator must not empty")
	}

	operator = fmt.Sprintf(" %s ", operator)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, p.TableName(), strings.Join(filter, operator))
	row, err := sqliteConn.Query(query)
	if err != nil {
		return Users{}, err
	}
	defer row.Close()

	var output Users
	for row.Next() {
		var currentData Users
		if err = row.Scan(
			&currentData.ID,
			&currentData.Email,
			&currentData.Username,
			&currentData.Phone,
			&currentData.Name,
			&currentData.Password,
			&currentData.RoleID,
			&currentData.CreatedAt,
		); err != nil {
			return Users{}, err
		}

		output = currentData
	}

	return output, nil
}
