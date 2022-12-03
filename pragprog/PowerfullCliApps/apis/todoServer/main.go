package main

import "github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/todo"

type todoResponse struct {
	Results todo.List `json:"Results"`
}
