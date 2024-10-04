package teststore

import (
	"kood-real-time-forum/internal/models"
	"kood-real-time-forum/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*models.User),
	}

	return s.userRepository
}
