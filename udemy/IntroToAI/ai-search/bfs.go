package main

import (
	"fmt"
	"log"
	"slices"
)

// BreadthFirstSearch struct implements the BFS algorithm
type BreadthFirstSearch struct {
	Frontier []*Node
	Game     *Maze
}

func (bfs *BreadthFirstSearch) GetFrontier() []*Node {
	return bfs.Frontier
}

func (bfs *BreadthFirstSearch) Add(node *Node) {
	bfs.Frontier = append(bfs.Frontier, node)
}

func (bfs *BreadthFirstSearch) ContainsState(i *Node) bool {
	for _, node := range bfs.Frontier {
		if node.State == i.State {
			return true
		}
	}
	return false
}

func (bfs *BreadthFirstSearch) Empty() bool {
	return len(bfs.Frontier) == 0
}

func (bfs *BreadthFirstSearch) Remove() (*Node, error) {
	if !bfs.Empty() {
		if bfs.Game.Debug {
			fmt.Println("BFS Frontier before remove:")
			for _, n := range bfs.Frontier {
				fmt.Printf("Node State: %v\n", n.State)
			}
		}
		// BFS uses FIFO (First In, First Out) - remove from the beginning
		node := bfs.Frontier[0]
		bfs.Frontier = bfs.Frontier[1:]
		return node, nil
	}
	return nil, fmt.Errorf("Frontier is empty")
}

func (bfs *BreadthFirstSearch) Solve() {
	fmt.Println("Solving maze using Breadth-First Search...")
	bfs.Game.NumExplored = 0

	start := Node{
		State:  bfs.Game.Start,
		Parent: nil,
		Action: "",
	}
	bfs.Add(&start)
	bfs.Game.CurrentNode = &start

	for {
		if bfs.Empty() {
			return
		}

		currentNode, err := bfs.Remove()
		if err != nil {
			log.Println("Error removing node from frontier:", err)
			return
		}

		if bfs.Game.Debug {
			fmt.Println("Removed:", currentNode.State)
			fmt.Println("---------------")
			fmt.Println()
		}

		bfs.Game.CurrentNode = currentNode
		bfs.Game.NumExplored += 1

		// Have we found the solution?
		if bfs.Game.Goal == currentNode.State {
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

			bfs.Game.Solution = Solution{
				Actions: actions,
				Cells:   cells,
			}

			bfs.Game.Explored = append(bfs.Game.Explored, currentNode.State)
			break
		}

		bfs.Game.Explored = append(bfs.Game.Explored, currentNode.State)

		// Build animation frame if appropriate
		if bfs.Game.Animate {
			bfs.Game.OutputImage(fmt.Sprintf("tmp/%06d.png", bfs.Game.NumExplored))
		}

		for _, neighbor := range bfs.Neighbors(currentNode) {
			if !bfs.ContainsState(neighbor) {
				if !inExplored(neighbor.State, bfs.Game.Explored) {
					bfs.Add(&Node{
						State:  neighbor.State,
						Parent: currentNode,
						Action: neighbor.Action,
					})
				}
			}
		}
	}
}

func (bfs *BreadthFirstSearch) Neighbors(node *Node) []*Node {
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
		if 0 <= candidate.State.Row && candidate.State.Row < bfs.Game.Height {
			if 0 <= candidate.State.Col && candidate.State.Col < bfs.Game.Width {
				if !bfs.Game.Walls[candidate.State.Row][candidate.State.Col].wall {
					neighbors = append(neighbors, candidate)
				}
			}
		}
	}

	// BFS does not randomize neighbors - explores in consistent order
	return neighbors
}
