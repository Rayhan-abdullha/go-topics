// repository/user_mysql_repository.go
package repository

import (
	"database/sql"
	"myapp/domain"
)

type UserMySQLRepository struct {
	DB *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{DB: db}
}

func (r *UserMySQLRepository) FindByID(id int) (*domain.User, error) {
	user := domain.User{}

	row := r.DB.QueryRow(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserMySQLRepository) Save(user *domain.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(name, email) VALUES(?, ?)",
		user.Name,
		user.Email,
	)
	return err
}
