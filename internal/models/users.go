package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             uint
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(email, password string) (uint, error) {
	return 0, nil
}

func (m *UserModel) Exist(id uint) (bool, error) {
	return false, nil
}
