package users

import (
	"fmt"
	"github.com/apm-dev/go-user-rest-api/datasources/mysql/users_db"
	"github.com/apm-dev/go-user-rest-api/utils/date_utils"
	"github.com/apm-dev/go-user-rest-api/utils/errors"
	"log"
	"strings"
)

var (
	queryGetAll       = "SELECT * FROM users"
	queryFindUserById = "SELECT * FROM users WHERE id=?"
	queryInsertUser   = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryUpdateUser   = "UPDATE users SET first_name=?, last_name=?, email=?, updated_at=? WHERE id = ?;"
	queryDeleteUser   = "DELETE FROM users WHERE id = ?"
)

func (User) All() ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetAll)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt,
			&user.UpdatedAt); err != nil {
			return nil, errors.InternalServerError(err.Error())
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFound("There is no record in our database")
	}
	return results, nil
}

func (u *User) Find() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryFindUserById)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt,
		&u.UpdatedAt); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "no rows in result") {
			return errors.NotFound(fmt.Sprintf("user %d not found", u.ID))
		}
		return errors.InternalServerError(fmt.Sprintf("failed when trying to get user % d", u.ID))
	}
	return nil
}

func (u *User) Insert() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("Error when trying to get last user_id: %s", err.Error()))
	}
	u.ID = userId
	return nil
}

func (u *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	updatedAt := date_utils.GetNow()
	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, updatedAt, u.ID)
	if err != nil {
		return errors.InternalServerError("Failed when trying to update user")
	}
	u.UpdatedAt = updatedAt
	return nil
}

func (u *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
