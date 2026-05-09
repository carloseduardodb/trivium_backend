package repository

import (
	"fmt"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/postgres/connection"
)

type UserRepository struct{}

func NewUserRepository() repositorier.UserRepositorier {
	return &UserRepository{}
}

func (r *UserRepository) Save(user entity.User) (entity.User, error) {
	db := connection.GetDB()

	query := `INSERT INTO users (name, email, photo_path, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query,
		user.Name,
		user.Email,
		user.PhotoPath,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) FindById(id int64) (*entity.User, error) {
	db := connection.GetDB()

	var user entity.User
	err := db.QueryRowx(`
		SELECT id, name, email, photo_path
		FROM users WHERE id = $1`, id).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	db := connection.GetDB()

	var user entity.User
	err := db.QueryRowx(`
		SELECT id, name, email, photo_path
		FROM users WHERE email = $1`, email).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	db := connection.GetDB()

	var users []entity.User
	err := db.Select(&users, "SELECT id, name, email, photo_path FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(user entity.User) (entity.User, error) {
	db := connection.GetDB()

	query := `UPDATE users SET name = $1, email = $2, photo_path = $3 WHERE id = $4`
	result, err := db.Exec(query, user.Name, user.Email, user.PhotoPath, user.ID)
	if err != nil {
		return entity.User{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}
	if rows == 0 {
		return entity.User{}, fmt.Errorf("user with id %d not found", user.ID)
	}

	return user, nil
}

func (r *UserRepository) Delete(id int64) error {
	db := connection.GetDB()

	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
