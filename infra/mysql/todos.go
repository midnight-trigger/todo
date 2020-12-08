package mysql

import (
	"fmt"
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
	FindById(id int64) (todo Todos, err error)
	Create(todo *Todos) (insertedTodo *Todos, err error)
	Update(oldParams Todos, updateParams map[string]interface{}) (err error)
}

func GetNewTodo() *Todos {
	return &Todos{}
}

func (m *Todos) FindById(id int64) (todo Todos, err error) {
	err = db.Where("id = ?", id).First(&todo).Error
	return
}

func (m *Todos) Create(todo *Todos) (insertedTodo *Todos, err error) {
	err = db.Create(todo).Error
	insertedTodo = todo
	return
}

func (m *Todos) Update(oldParams Todos, updateParams map[string]interface{}) (err error) {
	err = db.Model(&oldParams).Updates(updateParams).Error
	fmt.Println(oldParams)
	return
}
