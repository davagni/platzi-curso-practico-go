package main

import (
	"encoding/json"
	"net/http"
)

//Middleware a
type Middleware func(http.HandlerFunc) http.HandlerFunc

//User a
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//MetaData a
type MetaData interface{}

//ToJSON a
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
