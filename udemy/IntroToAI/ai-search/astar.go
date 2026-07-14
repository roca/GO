package main

import (
	"container/heap"
	"fmt"
	"log"
	"slices"
)

const FloodedCost = 100

// AstarSearch struct implements Dijkstra's algorithm
type AstarSearch struct {
	Frontier PriorityQueueAstar
	Game     *Maze
}

func (d *AstarSearch) GetFrontier() []*Node {
	return d.Frontier
}

func (d *AstarSearch) Add(node *Node) {
	node.CostToGoal = node.ManhattanDistance(d.Game.Start) // Cost from start to this node
	node.EstimatedCostToGoal = euclideanDist(node.State, d.Game.Goal) + float64(node.CostToGoal)
	if node.State.Water {
		node.EstimatedCostToGoal += FloodedCost
	}
	d.Frontier.Push(node)
	// Re-establish the heap property after adding a new node
	heap.Init(&d.Frontier)
}

func (d *AstarSearch) ContainsState(i *Node) bool {
	for _, node := range d.Frontier {
		if node.State == i.State {
			return true
		}
	}
	return false
}

func (d *AstarSearch) Empty() bool {
	return d.Frontier.Len() == 0
}

func (d *AstarSearch) Remove() (*Node, error) {
	if !d.Empty() {
		if d.Game.Debug {
			fmt.Println("Dijkstra Frontier before remove:")
			for _, n := range d.Frontier {
				fmt.Printf("Node State: %v, Cost: %d\n", n.State, n.CostToGoal)
			}
		}
		// Dijkstra uses a priority queue - remove node with lowest cost
		node := heap.Pop(&d.Frontier).(*Node)
		return node, nil
	}
	return nil, fmt.Errorf("Frontier is empty")
}

func (d *AstarSearch) Solve() {
	fmt.Println("Solving maze using A* Algorithm...")
	d.Game.NumExplored = 0

	// Initialize the priority queue
	d.Frontier = make(PriorityQueueAstar, 0)
	heap.Init(&d.Frontier)

	start := Node{
		State:      d.Game.Start,
		Parent:     nil,
		Action:     "",
		CostToGoal: 0, // Cost from start is 0
	}
	d.Add(&start)
	d.Game.CurrentNode = &start

	for {
		if d.Empty() {
			return
		}

		currentNode, err := d.Remove()
		if err != nil {
			log.Println("Error removing node from frontier:", err)
			return
		}

		if d.Game.Debug {
			fmt.Printf("Removed: %v (Cost: %d)\n", currentNode.State, currentNode.CostToGoal)
			fmt.Println("---------------")
			fmt.Println()
		}

		d.Game.CurrentNode = currentNode
		d.Game.NumExplored += 1

		// Have we found the solution?
		if d.Game.Goal == currentNode.State {
			var actions []string
			var cells []Point

			for {
				if currentNode.Parent != nil {
					actions = append(actions, currentNode.Action)
					cells = append(cells, currentNode.State)
					currentNode = currentNode.Parent
				} else {
					break
				}
			}

			// Reverse the actions and cells
			slices.Reverse(actions)
			slices.Reverse(cells)

			d.Game.Solution = Solution{
				Actions: actions,
				Cells:   cells,
			}

			d.Game.Explored = append(d.Game.Explored, currentNode.State)
			break
		}

		d.Game.Explored = append(d.Game.Explored, currentNode.State)

		// Build animation frame if appropriate
		if d.Game.Animate {
			d.Game.OutputImage(fmt.Sprintf("tmp/%06d.png", d.Game.NumExplored))
		}

		// for _, neighbor := range d.Neighbors(currentNode) {
		// 	if !d.ContainsState(neighbor) {
		// 		if !inExplored(neighbor.State, d.Game.Explored) {
		// 			d.Add(&Node{
		// 				State:  neighbor.State,
		// 				Parent: currentNode,
		// 				Action: neighbor.Action,
		// 			})
		// 		}
		// 	}
		// }
		for _, neighbor := range d.Neighbors(currentNode) {
			if !inExplored(neighbor.State, d.Game.Explored) {
				// Calculate the new cost (current cost + 1, since each move has cost 1)
				newCost := currentNode.CostToGoal + 1
				neighbor.CostToGoal = newCost

				// Check if this state is already in the frontier
				inFrontier := false
				for _, frontierNode := range d.Frontier {
					if frontierNode.State == neighbor.State {
						inFrontier = true
						// If we found a cheaper path, we should update it
						// But with uniform cost (all edges = 1), the first path is always cheapest
						break
					}
				}

				if !inFrontier {
					d.Add(&Node{
						State:      neighbor.State,
						Parent:     currentNode,
						Action:     neighbor.Action,
						CostToGoal: newCost,
					})
				}
			}
		}
	}
}

func (d *AstarSearch) Neighbors(node *Node) []*Node {
	row := node.State.Row
	col := node.State.Col

	candidates := []*Node{
		{State: Point{Row: row - 1, Col: col}, Parent: node, Action: "up"},
		{State: Point{Row: row + 1, Col: col}, Parent: node, Action: "down"},
		{State: Point{Row: row, Col: col - 1}, Parent: node, Action: "left"},
		{State: Point{Row: row, Col: col + 1}, Parent: node, Action: "right"},
	}

	var neighbors []*Node
	for _, candidate := range candidates {
		if 0 <= candidate.State.Row && candidate.State.Row < d.Game.Height {
			if 0 <= candidate.State.Col && candidate.State.Col < d.Game.Width {
				if !d.Game.Walls[candidate.State.Row][candidate.State.Col].wall {
					if d.Game.Walls[candidate.State.Row][candidate.State.Col].State.Water {
						candidate.State.Water = true
					}
					neighbors = append(neighbors, candidate)
				}
			}
		}
	}

	// Dijkstra does not randomize neighbors - explores in consistent order
	return neighbors
}
