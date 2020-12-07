package domain

import (
	"github.com/midnight-trigger/todo/infra/mysql"
)

type User struct {
	Base
	MUsers mysql.IUsers
}

func GetNewUserService() *User {
	user := new(User)
	user.MUsers = mysql.GetNewUser()
	return user
}
