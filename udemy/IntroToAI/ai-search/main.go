package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// IDs for search types
const (
	DFS      = iota // Depth-First Search
	BFS             // Breadth-First Search
	GBFS            // Greedy Best-First Search
	AStar           // A* Search
	DIJKSTRA        // Dijkstra's Algorithm
)

// Point represents a coordinate in the maze
type Point struct {
	Row int
	Col int
}

// Wall represents a cell in the maze, indicating if it's a wall or open space
type Wall struct {
	State Point
	wall  bool
}

// Node represents a state in the search tree,
// including its index, state, parent node, and action taken to reach it
type Node struct {
	index      int
	State      Point
	Parent     *Node
	Action     string
	CostToGoal int
}

func (n *Node) ManhattanDistance(goal Point) int {
	return abs(n.State.Row-goal.Row) + abs(n.State.Col-goal.Col)
}

// Solution represents the result of a search,
type Solution struct {
	Actions []string
	Cells   []Point
}

// Maze represents the maze structure with its dimensions, start and goal points, and wall layout
type Maze struct {
	Height      int
	Width       int
	Start       Point
	Goal        Point
	Walls       [][]Wall
	CurrentNode *Node
	Solution    Solution
	Explored    []Point
	Steps       int
	NumExplored int
	Debug       bool
	SearchType  int
	Animate     bool
}

func init() {
	_ = os.Mkdir("./tmp", os.ModePerm)
	// Ensure the tmp directory is empty at the start
	emptyTmp()
}

func main() {
	var m Maze
	var maze, searchType string

	flag.StringVar(&maze, "file", "maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type: dfs, bfs, gbfs, astar, dijkstra")
	flag.BoolVar(&m.Debug, "debug", false, "write debugging info to console")
	flag.BoolVar(&m.Animate, "animate", false, "produce animation of search process")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startTime := time.Now()

	switch strings.ToLower(searchType) {
	case "dfs":
		m.SearchType = DFS
		solveDFS(&m)
	case "bfs":
		m.SearchType = BFS
		solveBFS(&m)
	default:
		fmt.Printf("Search type %s not recognized\n", searchType)
		os.Exit(1)
	}

	if len(m.Solution.Actions) > 0 {
		fmt.Println("Solution:")
		// m.printMaze()
		fmt.Println("Solution is", len(m.Solution.Cells), "steps.")
		fmt.Println("Time to solve:", time.Since(startTime))
		m.OutputImage("image.png")
	} else {
		fmt.Println("No solution found.")
	}

	fmt.Println("Total nodes explored:", len(m.Explored))

	if m.Animate {
		fmt.Println("Generating animation...")
		m.OutPutAnimatedImage()
		fmt.Println("Animation complete: animation.png")
	}
}

func (g *Maze) printMaze() {
	for r, row := range g.Walls {
		for c, col := range row {
			if col.wall {

				// Get big square icon
				// Press control + command + space to access emoji menu
				fmt.Print("█")
			} else if g.Start.Row == col.State.Row && g.Start.Col == col.State.Col {
				fmt.Print("A")
			} else if g.Goal.Row == col.State.Row && g.Goal.Col == col.State.Col {
				fmt.Print("B")
			} else if g.inSolution(Point{Row: r, Col: c}) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (g *Maze) inSolution(p Point) bool {
	for _, step := range g.Solution.Cells {
		if step.Row == p.Row && step.Col == p.Col {
			return true
		}
	}
	return false
}

func solveDFS(m *Maze) {
	var s DepthFirstSearch
	s.Game = m
	fmt.Println("Goal is at:", s.Game.Goal)
	s.Solve()
}

func solveBFS(m *Maze) {
	var s BreadthFirstSearch
	s.Game = m
	fmt.Println("Goal is at:", s.Game.Goal)
	s.Solve()
}

func (g *Maze) Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", fileName, err)
	}
	defer f.Close()

	var fileContents []string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file %s: %v", fileName, err)
		}
		fileContents = append(fileContents, line)
	}

	foundStart, foundEnd := false, false
	for _, line := range fileContents {
		if strings.Contains(line, "A") {
			foundStart = true
		}
		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}

	if !foundStart {
		return fmt.Errorf("start point 'A' not found in the maze")
	}

	if !foundEnd {
		return fmt.Errorf("end point 'B' not found in the maze")
	}

	g.Height = len(fileContents)
	g.Width = len(fileContents[0])

	var rows [][]Wall

	for i, row := range fileContents {
		var cols []Wall
		for j, col := range row {
			curLetter := fmt.Sprintf("%c", col)
			var wall Wall
			switch curLetter {
			case "A":
				g.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "B":
				g.Goal = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = true
			default:
				continue
			}
			cols = append(cols, wall)
		}
		rows = append(rows, cols)
	}

	g.Walls = rows
	return nil
}
