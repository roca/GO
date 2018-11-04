// File name: ...\s15\08_reflection3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"reflect" 
	"strconv"
)

func main() {
	i1 := int64(10)
	fmt.Printf("%-10v %v \n", getType(i1), getType([]int64{i1}))

	i2 := int32(10)
	fmt.Printf("%-10v %v \n", getType(i2), getType([]int32{i2}))

	b := true
	fmt.Printf("%-10v %v \n", getType(b), getType([]bool{b}))

	s := "hello"
	fmt.Printf("%-10v %v \n", getType(s), getType([]string{s}))

	f := float32(10.31)
	fmt.Printf("%-10v %v \n", getType(f), getType([]float32{f}))
}

func getType(value interface{}) string {

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int32, reflect.Int64: // and other types ...
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Uint, reflect.Uint8: // and other types ...
		return strconv.FormatUint(v.Uint(), 10)

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 2, 64)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.String:
		return strconv.Quote(v.String())

	// other cases such as reflect.Map, ...

	default:
		return v.Type().String() + " value"
	}
}
