package api

type void interface{}

type UserRoutes struct {
	PostSignupUser void `method:"POST" path:"api/v1/users/signup"`
	PostSigninUser void `method:"POST" path:"api/v1/users/signin"`
}

type TodoRoutes struct {
	GetTodos   void `method:"GET"    path:"api/v1/todos"         auth:"true"`
	PostTodo   void `method:"POST"   path:"api/v1/todos"         auth:"true"`
	PutTodo    void `method:"PUT"    path:"api/v1/todos/:todoId" auth:"true"`
	PatchTodo  void `method:"PATCH"  path:"api/v1/todos/:todoId" auth:"true"`
	DeleteTodo void `method:"DELETE" path:"api/v1/todos/:todoId" auth:"true"`
}
