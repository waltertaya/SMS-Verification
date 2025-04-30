package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type User struct {
//
// 	ID       uuid.UUID `gorm:"primary_key" json:"id"`
// 	Email    string    `gorm:"unique" json:"email"`
// 	Username string    `gorm:"unique" json:"username"`
// 	Password string    `json:"password"`
// }

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username string    `gorm:"type:varchar(150);not null;uniqueIndex" json:"username"`
	Name     string    `gorm:"type:varchar(150);not null" json:"name"`
	Email    string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Role     string    `gorm:"type:varchar(100);not null" json:"role"`
	Verified bool      `gorm:"not null" json:"verified"`
}
