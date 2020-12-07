package api

type void interface{}

type UserRoutes struct {
	PostUser void `method:"POST" path:"api/v1/users"`
}
