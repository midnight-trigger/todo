package mysql

import (
	"time"
)

type Todos struct {
	Id        int64      `gorm:"type:bigint(20);primary_key;auto_increment"`
	UserId    string     `gorm:"type:char(36);not null"`
	Title     string     `gorm:"type:varchar(255);not null"`
	Body      string     `gorm:"type:text;not null"`
	Status    string     `gorm:"type:enum('todo','progress','finished');not null;default:'todo'"`
	CreatedAt time.Time  `gorm:"type:datetime;not null"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null"`
	DeletedAt *time.Time `gorm:"type:datetime"`
}

//go:generate mockgen -source todos.go -destination mock/mock_todos.go
type ITodos interface {
}

func GetNewTodo() *Todos {
	return &Todos{}
}
