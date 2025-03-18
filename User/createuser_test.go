package user

import (
	"crypto/sha256"
	"testing"
)

func TestCreateUser(t *testing.T) {
	username := "testuser"
	user := CreateUser(username)

	if user.Name != username {
		t.Errorf("Expected username %s, got %s", username, user.Name)
	}

	if len(user.Userid) == 0 {
		t.Error("UserID should not be empty")
	}
}

func TestCreateUserID(t *testing.T) {
	id := CreateUserID(8)

	if len(id) != sha256.Size {
		t.Errorf("Expected hash size %d, got %d", sha256.Size, len(id))
	}
}

func TestRandBytes(t *testing.T) {
	b, err := RandBytes(8)
	if err != nil {
		t.Errorf("randBytes returned an error: %v", err)
	}

	if len(b) != 8 {
		t.Errorf("Expected length 8, got %d", len(b))
	}
}
