// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ItemCatalog struct {
	ID          int     `json:"id"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
