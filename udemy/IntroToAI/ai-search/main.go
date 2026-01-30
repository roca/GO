package main

const (
	DFS = iota
	BFS
	GBFS
	AStar
	DIJKSTRA
)

type Point struct {
	Row int
	Col int
}

type Wall struct {
	State Point
	wall  bool
}
