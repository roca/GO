package main

import "os"

func inExplored(needle Point, haystack []Point) bool {
	for _, point := range haystack {
		if point.Row == needle.Row && point.Col == needle.Col {
			return true
		}
	}
	return false
}

func emptyTmp() {
	directory := "./tmp/"
	dir, _ := os.Open(directory)
	filesToDelete, _ := dir.Readdir(0)

	for index := range filesToDelete {
		f := filesToDelete[index]
		fullpath := directory + f.Name()
		_ = os.Remove(fullpath)
	}
}
