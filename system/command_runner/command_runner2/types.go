package main

import (	
	"runtime"
)

var workers = runtime.NumCPU()

type Group struct {
	Id int
}

type Result struct {
	output string
}

type Job struct {
	Command
	results chan<- Result
}

type Command struct {
	Id    int
	path  string
	dir   string
	async int
}