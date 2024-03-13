package mocks

import "lets-go-snippetbox/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (uint, error) {
	if email == "login@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id uint) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
