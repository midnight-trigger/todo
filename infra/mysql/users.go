package mysql

import (
	"time"
)

type Users struct {
	Id       string `gorm:"type:char(36);primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	// Email     string     `gorm:"type:varchar(255);unique_index;not null"`
	// Password  string     `gorm:"type:varchar(255);not null"`
	// IsVerify  bool       `gorm:"type:tinyint(1);not null"`
	CreatedAt time.Time  `gorm:"type:datetime;not null"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null"`
	DeletedAt *time.Time `gorm:"type:datetime"`
}

//go:generate mockgen -source users.go -destination mock/mock_users.go
type IUsers interface {
	Create(user *Users) (err error)
}

func GetNewUser() *Users {
	return &Users{}
}

func (m *Users) Create(user *Users) (err error) {
	err = db.Create(user).Error
	return
}
