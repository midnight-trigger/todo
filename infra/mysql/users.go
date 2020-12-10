package mysql

import (
	"time"
)

type Users struct {
	Id        string     `gorm:"type:char(36);primary_key"`
	Username  string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"type:datetime;not null"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null"`
	DeletedAt *time.Time `gorm:"type:datetime"`
}

//go:generate mockgen -source users.go -destination mock/mock_users.go
type IUsers interface {
	FindById(id string) (user Users, err error)
	Create(user *Users) (err error)
}

func GetNewUser() *Users {
	return &Users{}
}

func (m *Users) FindById(id string) (user Users, err error) {
	err = db.Where("id = ?", id).First(&user).Error
	return
}

func (m *Users) Create(user *Users) (err error) {
	err = db.Create(user).Error
	return
}
