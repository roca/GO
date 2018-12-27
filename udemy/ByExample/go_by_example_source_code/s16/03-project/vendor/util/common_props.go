package util

import (
	"sync"
)

// CPUUsed ... its value is set in the 'init()' function of the main.
var CPUUsed int

// WaitG ...
var WaitG sync.WaitGroup

// Mu ...
var Mu sync.Mutex

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		// panic(err)
		panic(err.Error())
	}
}
