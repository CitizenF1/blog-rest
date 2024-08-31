package memory

import (
	"blog-rest/internal/models"
	"errors"
	"sync"
	"time"
)

type BlogStorage struct {
	users map[int]models.User
	posts map[int]models.Post
	mu    sync.RWMutex
}

func NewBlogStorage() *BlogStorage {
	return &BlogStorage{
		users: make(map[int]models.User),
		posts: make(map[int]models.Post),
	}
}

func (s *BlogStorage) AddUser(user models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = len(s.users) + 1

	user.CreatedAt = time.Now()

	user.LastModifiedAt = user.CreatedAt

	s.users[user.ID] = user

	return nil
}

func (s *BlogStorage) GetUsers() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.User
	for _, user := range s.users {
		result = append(result, user)
	}

	return result
}

func (s *BlogStorage) UpdateUser(id int, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if user, exists := s.users[id]; exists {
		user.Name = name

		user.LastModifiedAt = time.Now()

		s.users[id] = user

		return nil
	}

	return errors.New("user not found")
}

func (s *BlogStorage) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; exists {

		delete(s.users, id)

		return nil
	}

	return errors.New("user not found")
}

func (s *BlogStorage) AddPost(post models.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post.ID = len(s.posts) + 1

	post.CreatedAt = time.Now()

	s.posts[post.ID] = post

	return nil
}
