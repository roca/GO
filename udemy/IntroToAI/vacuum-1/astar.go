package main

import (
	"container/heap"
	"math"
)

type PQItem struct {
	point    Point
	priority float64
	index    int
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq)}

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
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0:n-1]
	return item
}

func (pq *PriorityQueue) Update(item *PQItem, priority float64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Astar(room *Room, start, goal Point) []Point {
	if !room.IsValid(start.X, start.Y) || !room.IsValid(goal.X, goal.Y) {
		return []Point{}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	openSetItems := make(map[Point]*PQItem)

	closedSet := make(map[Point]bool)
	gScore := make(map[Point]float64)
	fScore := make(map[Point]float64)
	cameFrom := make(map[Point]Point)

	gScore[start] = 0
	fScore[start] = heuristic(start, goal)

	startItem := &PQItem{
		point: start,
		priority: fScore[start],
		index: 0,
	}
	heap.Push(&pq, startItem)
	openSetItems[start] = startItem

	// Main A* loop
	for pq.Len() > 0 {
		// Get the point with the lowest f-score from priority queue
		currentItem := heap.Pop(&pq).(*PQItem)
		current := currentItem.point
		delete(openSetItems, current)

		// If we've reached goal, reconstruct and return the path
		if current.X == goal.X && current.Y == goal.Y {
			return reconstructPath(cameFrom, current)
		}

		// Mark current point as processed.
		closedSet[current] = true

		// Check all neighbors
		for _, dir := range directions {
			neighbor := Point{X: current.X + dir[0], Y: current.Y + dir[1]}

			// Skip if neighbor is invalid or in the closed set
			if !room.IsValid(neighbor.X, neighbor.Y) || closedSet[neighbor] {
				continue
			}

			// calculate tentative g-score
			tentativeGScore := gScore[current] + 1

			if _, exists := gScore[neighbor]; !exists || tentativeGScore < gScore[neighbor] {
				// Updat path information
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + heuristic(neighbor, goal)

				// update priority queue
				if item, exists := openSetItems[neighbor]; exists {
					pq.Update(item, fScore[neighbor])
				} else {
					// Add new point to the priorty queue
					neigborItem := &PQItem{
						point: neighbor,
						priority: fScore[neighbor],
					}
					heap.Push(&pq, neigborItem)
					openSetItems[neighbor] = neigborItem
				}
			}
		}
	}

	return []Point{}
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