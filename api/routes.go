package api

type void interface{}

type UserRoutes struct {
	// 会員登録
	PostSignupUser void `method:"POST" path:"api/v1/users/signup"`
	// ログイン
	PostSigninUser void `method:"POST" path:"api/v1/users/signin"`
}

type TodoRoutes struct {
	// Todo検索・一覧取得
	GetTodos void `method:"GET" path:"api/v1/todos" auth:"true"`
	// Todo新規作成
	PostTodo void `method:"POST" path:"api/v1/todos" auth:"true"`
	// Todo内容更新
	PutTodo void `method:"PUT" path:"api/v1/todos/:todoId" auth:"true"`
	// Todoステータス更新
	PatchTodo void `method:"PATCH" path:"api/v1/todos/:todoId" auth:"true"`
	// Todo削除
	DeleteTodo void `method:"DELETE" path:"api/v1/todos/:todoId" auth:"true"`
}
