package repository

import (
	"database/sql"
	"fmt"
	"github.com/faneaatiku/auth_api/app/entity"
	"github.com/labstack/echo/v4"
)

type UserRepository struct {
	logger echo.Logger
	mysql  *sql.DB
}

func NewUserRepository(mysql *sql.DB, logger echo.Logger) (*UserRepository, error) {
	if mysql == nil {
		return nil, fmt.Errorf("could not create user repository: invalid dependencies")
	}

	return &UserRepository{mysql: mysql, logger: logger}, nil
}

func (r UserRepository) GetUserByEmail(email string) (entity.User, bool) {
	var user entity.User
	stmtOut, err := r.mysql.Prepare("SELECT * FROM user as u WHERE u.canonical_email = ?")
	if err != nil {
		r.logger.Errorf("error on preparing statement: %v", err)

		return user, false
	}

	defer stmtOut.Close()
	rows := stmtOut.QueryRow(email)
	err = rows.Scan(&user)
	if err != nil {
		r.logger.Errorf("error on scanning query result: %v", err)

		return user, false
	}

	return user, true
}
