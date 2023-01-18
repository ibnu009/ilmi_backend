package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"html"
	"ilmi_backend/auth"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;unique;" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Token    string `gorm:"size:255;not null;" json:"token"`
	Otp      uint32 `gorm:"not null;" json:"otp"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (u *User) PrepareUser() {
	u.Id = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
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
		if u.Name == "" {
			return errors.New("Required Name")
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
func (u *User) GetUserProfile(db *gorm.DB, id uint32) (*User, error) {
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) IsUserExist(db *gorm.DB) bool {
	if err := db.Where("email = ?", u.Email).Take(&u).Error; err != nil {
		return false
	}
	return true
}

// chek email Otp
func (u *User) ChekEmail(db *gorm.DB, email string) (*User, error) {
	err := db.Where("email = ?", email).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateOtp(db *gorm.DB, email string, otp uint32) error {
	// Update with conditions
	err := db.Model(&User{}).Where("email = ?", email).Take(&User{}).Update("otp", otp).Error
	if err != nil {
		return err
	}
	return nil
}

// cek OTP
func (u *User) ChekOtp(db *gorm.DB, c *gin.Context, otp uint64) (*User, error) {
	if err := db.Debug().Model(&User{}).Where("otp = ?", otp).Take(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

// chek email if there email auto create token
func (u *User) ChekEmailOauth2(db *gorm.DB, email string) (string, error) {
	err := db.Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}
	return auth.CreateToken(u.Id)
}

// update password
func (u *User) ResertPassword(db *gorm.DB, c *gin.Context, password string) (*User, error) {
	if err := db.Model(&User{}).Where("id = ?", c.Param("id")).Take(&User{}).Update("password", password).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
