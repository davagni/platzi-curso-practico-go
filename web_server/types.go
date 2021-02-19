package main

import "net/http"

//Middleware a
type Middleware func(http.HandlerFunc) http.HandlerFunc

//User a
type User struct {
	Name  string
	Email string
	Phone string
}

//MetaData a
type MetaData interface{}
