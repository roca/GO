package main

import "fmt"

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
