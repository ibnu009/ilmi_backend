package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nama     string `gorm:"size:255;not null;" json:"nama"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Token    string `gorm:"size:255;not null;" json:"token"`
}

func (u *User) PrepareUser() {
	u.Id = 0
	u.Nama = html.EscapeString(strings.TrimSpace(u.Nama))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Token = html.EscapeString(strings.TrimSpace(u.Token))
}

func HashPasswordToSha256(password string) string {
	sum := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(sum[:])
	return hashedPassword
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nama == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// Create
func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}

// Read
func (u *User) GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	err := db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Update

// Delete
