package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name" validate:"required,min=3,max=50"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	PasswordHash string    `json:"-"`
	Monitors     []Monitor `gorm:"foreignKey:UserID" json:"monitors,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) SetPassword(raw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing raw password")
		return err
	}

	u.PasswordHash = string(hash)

	return nil
}

func (u *User) CheckPassword(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(raw)) == nil
}
