package teststore

import (
	"errors"

	"forum/internal/models"
)

type UserRepository struct {
	store *Store
	users map[string]*models.User
}

func (r *UserRepository) Create(user *models.User) error {
	// if err := user.BeforeCreate(); err != nil {
	// 	return err
	// }

	r.users[user.ID] = user

	return nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("not found")
}

func (r *UserRepository) Check(login string) (*models.User, error) {
	panic("need implement")
}

func (r *UserRepository) Delete(str string) error
func (r *UserRepository) GetAllOtherUsers(user_id string) ([]models.User, error)
func (r *UserRepository) IsUser(id string) (bool, error)
