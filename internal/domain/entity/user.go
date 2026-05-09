package entity

import (
	"fmt"
	"regexp"
)

type User struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	PhotoPath string `json:"photo_path" db:"photo_path"`
}

func isValidEmail(email string) bool {
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	rgx, err := regexp.Compile(re)
	if err != nil {
		fmt.Println("Erro na expressão regular:", err)
		return false
	}
	return rgx.MatchString(email)
}

func NewUser(name string, email string, photoPath string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if email == "" || !isValidEmail(email) {
		return nil, fmt.Errorf("email is required")
	}

	if photoPath == "" {
		photoPath = ""
	}

	return &User{
		Name:      name,
		Email:     email,
		PhotoPath: photoPath,
	}, nil
}
