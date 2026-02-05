package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
)

// DepthFirstSearch struct implements the DFS algorithm
type DepthFirstSearch struct {
	Frontier []*Node
	Game     *Maze
}

func (dfs *DepthFirstSearch) GetFrontier() []*Node {
	return dfs.Frontier
}

func (dfs *DepthFirstSearch) Add(node *Node) {
	dfs.Frontier = append(dfs.Frontier, node)
}

func (dfs *DepthFirstSearch) ContainsState(i *Node) bool {
	for _, node := range dfs.Frontier {
		if node.State == i.State {
			return true
		}
	}
	return false
}

func (dfs *DepthFirstSearch) Empty() bool {
	return len(dfs.Frontier) == 0
}

func (dfs *DepthFirstSearch) Remove() (*Node, error) {
	if !dfs.Empty() {
		if dfs.Game.Debug {
			fmt.Println("DFS Frontier before remove:")
			for _, n := range dfs.Frontier {
				fmt.Printf("Node State: %v\n", n.State)
			}
		}
		node := dfs.Frontier[len(dfs.Frontier)-1]
		dfs.Frontier = dfs.Frontier[:len(dfs.Frontier)-1]
		return node, nil
	}
	return nil, fmt.Errorf("Frontier is empty")
}

func (dfs *DepthFirstSearch) Solve() {
	fmt.Println("Solving maze using Depth-First Search...")
	dfs.Game.NumExplored = 0

	start := Node{
		State:  dfs.Game.Start,
		Parent: nil,
		Action: "",
	}
	dfs.Add(&start)
	dfs.Game.CurrentNode = &start

	for {
		if dfs.Empty() {
			return
		}

		currentNode, err := dfs.Remove()
		if err != nil {
			log.Println("Error removing node from frontier:", err)
			return
		}

		if dfs.Game.Debug {
			fmt.Println("Removed:", currentNode.State)
			fmt.Println("---------------")
			fmt.Println()
		}

		dfs.Game.CurrentNode = currentNode
		dfs.Game.NumExplored += 1

		// Have we found the solution?
		if dfs.Game.Goal == currentNode.State {
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

			dfs.Game.Solution = Solution{
				Actions: actions,
				Cells:   cells,
			}

			dfs.Game.Explored = append(dfs.Game.Explored, currentNode.State)
			break
		}

		dfs.Game.Explored = append(dfs.Game.Explored, currentNode.State)

		// Build animation frame if appropriate
		if dfs.Game.Animate {
			dfs.Game.OutputImage(fmt.Sprintf("tmp/%06d.png", dfs.Game.NumExplored))
		}

		for _, neighbor := range dfs.Neighbors(currentNode) {
			if !dfs.ContainsState(neighbor) {
				if !inExplored(neighbor.State, dfs.Game.Explored) {
					dfs.Add(&Node{
						State:  neighbor.State,
						Parent: currentNode,
						Action: neighbor.Action,
					})
				}
			}
		}
	}
}

func (dfs *DepthFirstSearch) Neighbors(node *Node) []*Node {
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
		if 0 <= candidate.State.Row && candidate.State.Row < dfs.Game.Height {
			if 0 <= candidate.State.Col && candidate.State.Col < dfs.Game.Width {
				if !dfs.Game.Walls[candidate.State.Row][candidate.State.Col].wall {
					neighbors = append(neighbors, candidate)
				}
			}
		}
	}

	for i := range neighbors {
		j := rand.Intn(i + 1)
		neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
	}

	return neighbors
}
