package main

import (
	"container/heap"
	"fmt"
	"log"
	"slices"
)

// GBFSSearch struct implements Greedy Best-First Search algorithm
type GBFSSearch struct {
	Frontier PriorityQueueGBFS
	Game     *Maze
}

func (g *GBFSSearch) GetFrontier() []*Node {
	return g.Frontier
}

func (g *GBFSSearch) Add(node *Node) {
	// GBFS uses only heuristic (Manhattan distance to goal)
	node.CostToGoal = node.ManhattanDistance(g.Game.Goal) // Heuristic cost from this node to goal
	g.Frontier.Push(node)
	// Re-establish the heap property after adding a new node
	heap.Init(&g.Frontier)
}

func (g *GBFSSearch) ContainsState(i *Node) bool {
	for _, node := range g.Frontier {
		if node.State == i.State {
			return true
		}
	}
	return false
}

func (g *GBFSSearch) Empty() bool {
	return g.Frontier.Len() == 0
}

func (g *GBFSSearch) Remove() (*Node, error) {
	if !g.Empty() {
		if g.Game.Debug {
			fmt.Println("GBFS Frontier before remove:")
			for _, n := range g.Frontier {
				fmt.Printf("Node State: %v, Heuristic: %d\n", n.State, n.CostToGoal)
			}
		}
		// GBFS uses a priority queue - remove node with lowest heuristic value
		node := heap.Pop(&g.Frontier).(*Node)
		return node, nil
	}
	return nil, fmt.Errorf("Frontier is empty")
}

func (g *GBFSSearch) Solve() {
	fmt.Println("Solving maze using Greedy Best-First Search...")
	g.Game.NumExplored = 0

	// Initialize the priority queue
	g.Frontier = make(PriorityQueueGBFS, 0)
	heap.Init(&g.Frontier)

	start := Node{
		State:      g.Game.Start,
		Parent:     nil,
		Action:     "",
		CostToGoal: 0, // Will be set by Add()
	}
	g.Add(&start)
	g.Game.CurrentNode = &start

	for {
		if g.Empty() {
			return
		}

		currentNode, err := g.Remove()
		if err != nil {
			log.Println("Error removing node from frontier:", err)
			return
		}

		if g.Game.Debug {
			fmt.Printf("Removed: %v (Heuristic: %d)\n", currentNode.State, currentNode.CostToGoal)
			fmt.Println("---------------")
			fmt.Println()
		}

		g.Game.CurrentNode = currentNode
		g.Game.NumExplored += 1

		// Have we found the solution?
		if g.Game.Goal == currentNode.State {
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

			g.Game.Solution = Solution{
				Actions: actions,
				Cells:   cells,
			}

			g.Game.Explored = append(g.Game.Explored, currentNode.State)
			break
		}

		g.Game.Explored = append(g.Game.Explored, currentNode.State)

		// Build animation frame if appropriate
		if g.Game.Animate {
			g.Game.OutputImage(fmt.Sprintf("tmp/%06d.png", g.Game.NumExplored))
		}

		// Explore neighbors - GBFS uses only heuristic, not cumulative cost
		for _, neighbor := range g.Neighbors(currentNode) {
			if !inExplored(neighbor.State, g.Game.Explored) && !g.ContainsState(neighbor) {
				g.Add(&Node{
					State:      neighbor.State,
					Parent:     currentNode,
					Action:     neighbor.Action,
					CostToGoal: 0, // Will be set by Add() using heuristic
				})
			}
		}
	}
}

func (g *GBFSSearch) Neighbors(node *Node) []*Node {
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
		if 0 <= candidate.State.Row && candidate.State.Row < g.Game.Height {
			if 0 <= candidate.State.Col && candidate.State.Col < g.Game.Width {
				if !g.Game.Walls[candidate.State.Row][candidate.State.Col].wall {
					neighbors = append(neighbors, candidate)
				}
			}
		}
	}

	// GBFS does not randomize neighbors - explores based on heuristic priority
	return neighbors
}

