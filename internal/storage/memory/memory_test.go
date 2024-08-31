package memory_test

import (
	"blog-rest/internal/models"
	"blog-rest/internal/storage/memory"
	"testing"
)

func TestCreateUser(t *testing.T) {
	store := memory.NewBlogStorage()
	user := models.User{
		ID:   1,
		Name: "John",
	}

	err := store.AddUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(store.GetUsers()) != 1 {
		t.Fatalf("expected 1 user, got %d", len(store.GetUsers()))
	}

	if store.GetUsers()[0].Name != "John" {
		t.Fatalf("expected user name to be 'John', got %s", store.GetUsers()[0].Name)
	}
}

func TestUpdateUser(t *testing.T) {
	store := memory.NewBlogStorage()
	user := models.User{
		ID:   1,
		Name: "John Doe",
	}

	_ = store.AddUser(user)

	err := store.UpdateUser(1, "Jane")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if store.GetUsers()[0].Name != "Jane" {
		t.Fatalf("expected user name to be 'Jane', got %s", store.GetUsers()[0].Name)
	}
}
