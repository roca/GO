package main

import (
	"container/heap"
	"math"
)

type PGItem struct {
	point    Point
	priority float64
	index    int
}

type PriorityQueue []*PGItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PGItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *PGItem, priority float64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Astar(room *Room, start Point, goal Point) []Point {
	if !room.IsValid(start.X, start.Y) || !room.IsValid(goal.X, goal.Y) {
		return []Point{}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	openSetItems := make(map[Point]*PGItem)

	closedSet := make(map[Point]bool)
	gScore := make(map[Point]float64)
	fScore := make(map[Point]float64)
	cameFrom := make(map[Point]Point)

	gScore[start] = 0
	fScore[start] = heuristic(start, goal)

	startItem := &PGItem{point: start, priority: fScore[start], index: 0}
	heap.Push(&pq, startItem)
	openSetItems[start] = startItem

	// Main A* loop
	for pq.Len() > 0 {
		//Get the point with the lowest f-score from the priority queue
		currentItem := heap.Pop(&pq).(*PGItem)
		current := currentItem.point
		delete(openSetItems, current)

		//If we've reached the goal, reconstruct the path and return it
		if current.X == goal.X && current.Y == goal.Y {
			return reconstructPath(cameFrom, current)
		}

		//Mark current point as processed by adding it to the closed set
		closedSet[current] = true

		// Check all neighbors of the current point
		for _, dir := range directions {
			neighbor := Point{X: current.X + dir[0], Y: current.Y + dir[1]}

			// Skip if neighbor is invalid or in the closed set
			if !room.IsValid(neighbor.X, neighbor.Y) || closedSet[neighbor] {
				continue
			}

			// Calculate tentative g-score for the neighbor
			tentativeGScore := gScore[current] + 1

			// If this path to neighbor is better, update scores and path
			if _, exists := gScore[neighbor]; !exists || tentativeGScore < gScore[neighbor] {
				// Update path information
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + heuristic(neighbor, goal)

				// Update priority queue
				if item, exists := openSetItems[neighbor]; exists {
					pq.Update(item, fScore[neighbor])
				} else {
					// Add new neighbor to the priority queue
					neighborItem := &PGItem{point: neighbor, priority: fScore[neighbor]}
					heap.Push(&pq, neighborItem)
					openSetItems[neighbor] = neighborItem
				}

			}
		}
	}

	return []Point{} // Return empty path if no path is found
}

func heuristic(a, b Point) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

func reconstructPath(cameFrom map[Point]Point, current Point) []Point {
	path := []Point{current}
	for {
		prev, exists := cameFrom[current]
		if !exists {
			break
		}
		path = append([]Point{prev}, path...)
		current = prev
	}
	return path
}
