package main

import "net/http"

//Middleware a
type Middleware func(http.HandlerFunc) http.HandlerFunc
