package main

import (
	"encoding/json"
)

// FunctionList ...
type FunctionList struct {
	Functions []Function
}

// Function ...
type Function struct {
	FunctionName string
}

// NewFunctionList ...
func NewFunctionList() (FunctionList, error) {
	dt, err := run("aws", "lambda", "list-functions")
	if err != nil {
		return FunctionList{}, err
	}

	var res FunctionList
	err = json.Unmarshal(dt, &res)

	return res, err
}

// HasFunction ...
func (fl FunctionList) HasFunction(fname string) bool {
	for _, v := range fl.Functions {
		if v.FunctionName == fname {
			return true
		}
	}
	return false
}
