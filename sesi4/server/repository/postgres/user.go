package postgres

import (
	"database/sql"
	"sesi4/server/model"
	"sesi4/server/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetUsers() (*[]model.User, error) {
	query := `
		SELECT 
			id, nip, name, address, 
			email, created_at, updated_at
		FROM
			employees 
	`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id, &user.Nip, &user.Fullname, &user.Address,
			&user.Email, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	rows.Close()

	return &users, nil
}

func (u *userRepo) Register(user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `
		INSERT INTO employees (
			id, nip, name, address, email, password, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Id, user.Nip, user.Fullname, user.Address,
		user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) FindUserByEmail(email string) (*model.User, error) {
	query := `
		SELECT 
			id, nip, name, address, password,
			email, created_at, updated_at
		FROM
			employees 
		WHERE email=$1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)
	var user model.User

	err = row.Scan(
		&user.Id, &user.Nip, &user.Fullname, &user.Address,
		&user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
