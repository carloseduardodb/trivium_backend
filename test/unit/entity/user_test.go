package entity_test

import (
	"testing"

	"trivium/internal/domain/entity"
)

func TestNewUser_Success(t *testing.T) {
	user, err := entity.NewUser("Carlos", "carlos@example.com", "https://photo.url/pic.jpg")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "Carlos" {
		t.Errorf("expected name 'Carlos', got '%s'", user.Name)
	}
	if user.Email != "carlos@example.com" {
		t.Errorf("expected email 'carlos@example.com', got '%s'", user.Email)
	}
	if user.PhotoPath != "https://photo.url/pic.jpg" {
		t.Errorf("expected photo path, got '%s'", user.PhotoPath)
	}
}

func TestNewUser_EmptyName(t *testing.T) {
	_, err := entity.NewUser("", "carlos@example.com", "photo.jpg")
	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestNewUser_InvalidEmail(t *testing.T) {
	_, err := entity.NewUser("Carlos", "invalid-email", "photo.jpg")
	if err == nil {
		t.Fatal("expected error for invalid email")
	}
}

func TestNewUser_EmptyEmail(t *testing.T) {
	_, err := entity.NewUser("Carlos", "", "photo.jpg")
	if err == nil {
		t.Fatal("expected error for empty email")
	}
}
