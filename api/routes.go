package api

type void interface{}

type UserRoutes struct {
	PostSigninUser void `method:"POST" path:"api/v1/users/signin"`
	PostUser       void `method:"POST" path:"api/v1/users"`
}

type TodoRoutes struct {
	PostTodo void `method:"POST" path:"api/v1/todos" auth:"true"`
}
