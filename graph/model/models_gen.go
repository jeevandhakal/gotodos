// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CompletedstatusInput struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	User        *User  `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
