package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

type Curve struct {
	DataPoints []Point
}

func (c *Curve) loadDataPointsFromFile(filePath string) (int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	lineCount := 0
	for e == nil {
		lineCount++
		data := strings.Replace(s, " ", "", -1)
		point := strings.Split(data, ",")
		pointX, errX := strconv.ParseFloat(point[0], 64)
		if errX != nil {
			//continue
		}
		pointY, errY := strconv.ParseFloat(point[1], 64)
		if errY != nil {
			//continue
		}
		fmt.Printf("point %d (x,y): (%g, %g)\n", len(c.DataPoints)+1, pointX, pointY)
		c.DataPoints = append(c.DataPoints, Point{pointX, pointY})
		s, e = Readln(r)
	}
	if len(c.DataPoints) != lineCount {
		return len(c.DataPoints), MyError{"Could not read some points"}
	}
	return len(c.DataPoints), nil
}
