# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains code for a Udemy course on Introduction to AI and Machine Learning with Go. The primary project is `ai-search`, which implements various pathfinding algorithms to solve maze problems.

## Project Structure

The repository uses Go workspaces (`go.work`) with a single module:
- `ai-search/` - Main AI search algorithms project for maze solving

## Commands

### Running the AI Search Program

```bash
cd ai-search
go run main.go -file mazes/maze.txt -search dfs
```

Available command-line flags:
- `-file <path>`: Maze file to solve (default: "maze.txt")
- `-search <algorithm>`: Search algorithm to use (default: "dfs")
- `-debug`: Enable verbose debugging output to console
- `-animate`: Generate an animated PNG showing the search process

Example with animation:
```bash
go run main.go -file mazes/maze-100-steps.txt -search dfs -animate
```

Search algorithms:
- `dfs` - Depth-First Search (implemented) - Uses LIFO stack, randomizes neighbor order
- `bfs` - Breadth-First Search (implemented) - Uses FIFO queue, guarantees shortest path
- `dijkstra` - Dijkstra's Algorithm (implemented) - Uses priority queue, guarantees shortest path
- `gbfs` - Greedy Best-First Search (implemented) - Uses priority queue with heuristic only
- `astar` - A* Search (implemented) - Uses priority queue with f(n) = g(n) + h(n), guarantees shortest path

### Building

```bash
cd ai-search
go build
```

### Running directly

```bash
cd ai-search
./ai-search -file mazes/maze.txt -search dfs
```

## Architecture

### Core Data Structures

All defined in `ai-search/main.go`:

- **Maze**: The main game state containing:
  - `Height`, `Width`: Maze dimensions
  - `Start`, `Goal`: Points marking the start ('A') and end ('B') positions
  - `Walls`: 2D array of Wall structs representing the maze grid
  - `CurrentNode`: Current search position
  - `Solution`: Final path from start to goal
  - `Explored`: Slice of Points that have been visited
  - `Steps`: Count of steps taken
  - `NumExplored`: Count of explored nodes
  - `Debug`: Boolean flag for verbose debug output
  - `SearchType`: Algorithm constant (DFS, BFS, etc.)
  - `Animate`: Boolean flag for animation generation

- **Point**: Row/column coordinate (used for positions in the maze)

- **Wall**: Cell in the maze grid with:
  - `State`: Point representing position
  - `wall`: boolean - `true` means impassable wall, `false` means open path

- **Node**: Search tree node containing:
  - `State`: Point representing position
  - `Parent`: Pointer to parent node (for path reconstruction)
  - `Action`: String describing how we got here ("up", "down", "left", "right")
  - `CostToGoal`: Integer cost (g(n) - actual cost from start, used by Dijkstra and A*)
  - `EstimatedCostToGoal`: Float64 estimated total cost (f(n) = g(n) + h(n), used by A*)
  - `index`: Integer used by priority queue implementations

- **Solution**: Final result with:
  - `Actions`: Slice of direction strings
  - `Cells`: Slice of Points representing the path

### Maze File Format

Maze files are stored in `ai-search/mazes/` as text files with:
- `A` - Start position
- `B` - Goal position
- `#` - Walls
- ` ` (space) - Open path
- Newlines define rows

Example:
```
##B   #
## ## #
#  #  #
# ## ##
     ##
A######
```

### Search Algorithm Pattern

Each search algorithm follows a common structure (see `dfs.go` for the DFS implementation):

1. **Algorithm Struct**: Contains:
   - `Frontier`: Slice of `*Node` representing the search frontier (or `PriorityQueueDijkstra` for Dijkstra, `PriorityQueueGBFS` for GBFS, `PriorityQueueAstar` for A*)
   - `Game`: Pointer to the `Maze` being solved

2. **Required Methods**:
   - `Add(node *Node)`: Add a node to the frontier
   - `Remove() (*Node, error)`: Remove and return a node from the frontier (LIFO for DFS, FIFO for BFS, priority for Dijkstra/GBFS)
   - `ContainsState(node *Node) bool`: Check if a state is already in the frontier
   - `Empty() bool`: Check if the frontier is empty
   - `Neighbors(node *Node) []*Node`: Generate valid neighboring states from current node
   - `Solve()`: Main search loop that explores the maze

3. **Search Logic Flow**:
   - Initialize frontier with start node
   - Loop: remove node from frontier, check if goal, add unexplored neighbors
   - Track explored states to avoid revisiting
   - Build solution by backtracking through parent pointers when goal is found

4. **IMPORTANT**: When checking if a cell is traversable in `Neighbors()`, use `wall == false` (open path), not `wall == true` (wall)

5. **Neighbor Ordering**:
   - DFS (`dfs.go`): Randomizes neighbor order before adding them to the frontier. Multiple runs may explore different paths even with the same maze.
   - BFS (`bfs.go`): Does NOT randomize - explores neighbors in consistent order (up, down, left, right) to guarantee shortest path.
   - Dijkstra (`dijkstra.go`): Does NOT randomize - uses priority queue based on `CostToGoal` (g(n)).
   - GBFS (`gbfs.go`): Does NOT randomize - uses priority queue based on heuristic (Manhattan distance to goal).
   - A* (`astar.go`): Does NOT randomize - uses priority queue based on f(n) = g(n) + h(n), where h(n) is Euclidean distance to goal.

6. **Algorithm Differences**:
   - **DFS**: LIFO (Last In, First Out) - removes from end of frontier like a stack
   - **BFS**: FIFO (First In, First Out) - removes from beginning of frontier like a queue
   - **Dijkstra**: Priority queue ordered by cumulative cost from start (g(n)) - removes node with lowest cost
   - **GBFS**: Priority queue ordered by heuristic only (h(n)) - removes node with lowest heuristic value (Manhattan distance to goal)
   - **A***: Priority queue ordered by f(n) = g(n) + h(n), where g(n) is actual cost from start and h(n) is Euclidean distance to goal - removes node with lowest estimated total cost
   - **Optimality**:
     - BFS guarantees shortest path but may explore more nodes (609 nodes for maze-100-steps.txt)
     - DFS may find a solution faster but doesn't guarantee optimality
     - Dijkstra guarantees shortest path and is optimal for weighted graphs (100 steps for maze-100-steps.txt)
     - GBFS is fast and explores fewer nodes (158 nodes for maze-100-steps.txt) but doesn't guarantee shortest path (108 steps)
     - A* guarantees shortest path (when using admissible heuristic) and is generally more efficient than Dijkstra

7. **Helper Functions**: `helpers.go` contains utility functions:
   - `inExplored(Point, []Point) bool`: Check if a Point has been visited
   - `emptyTmp()`: Clear the tmp directory of animation frames
   - `abs(int) int`: Return absolute value of an integer
   - `euclideanDist(Point, Point) float64`: Calculate Euclidean distance between two points (used by A* heuristic)

8. **Priority Queue Implementations**:
   - `priority-queue-dijkstra.go`: Priority queue for Dijkstra's algorithm (orders by cumulative cost from start, g(n))
   - `priority-queue-gbfs.go`: Priority queue for GBFS (orders by heuristic distance to goal, h(n))
   - `priority-queue-astar.go`: Priority queue for A* (orders by estimated total cost, f(n) = g(n) + h(n))
   - All implement Go's `heap.Interface` with `Len()`, `Less()`, `Swap()`, `Push()`, and `Pop()` methods

9. **Main Entry Point**: `main.go` handles file loading, validation, and dispatches to the appropriate solve function based on the `-search` flag
   - `solveDFS(m *Maze)`: Entry point for DFS
   - `solveBFS(m *Maze)`: Entry point for BFS
   - `solveGBFS(m *Maze)`: Entry point for GBFS
   - `solveDijkstra(m *Maze)`: Entry point for Dijkstra
   - `solveAStar(m *Maze)`: Entry point for A*

### Implementing a New Search Algorithm

To add a new search algorithm, follow this pattern:

1. **Create priority queue** (if using priority-based search): `priority-queue-<name>.go`
   - Implement `heap.Interface`: `Len()`, `Less()`, `Swap()`, `Push()`, `Pop()`
   - Define how nodes are ordered (see existing implementations for reference)

2. **Create algorithm file**: `<name>.go` with:
   - Algorithm struct with `Frontier` and `Game *Maze` fields
   - `Add(node *Node)`: Calculate and set node costs, add to frontier, call `heap.Init()` if priority queue
   - `Remove() (*Node, error)`: Remove node from frontier (with debug output if `Game.Debug`)
   - `ContainsState(node *Node) bool`: Check if state exists in frontier
   - `Empty() bool`: Check if frontier is empty
   - `Neighbors(node *Node) []*Node`: Generate valid neighbors (checking bounds and `!wall`)
   - `Solve()`: Main loop (initialize, explore until empty or goal found, track explored, build animation frames)
   - `solve<Name>(m *Maze)`: Entry point function

3. **Update `main.go`**:
   - Add constant to search type enum (line 13-19)
   - Add case to switch statement dispatching to solve function (line 95-108)
   - Update help text if needed

4. **Update CLAUDE.md**: Document the new algorithm in the search algorithms list and algorithm differences section

### Image Generation

The program generates visual output using `image.go`:

- **Static Image**: After solving, generates `image.png` showing the final solution with:
  - Black: Walls
  - Dark Green: Start position ('A')
  - Red: Goal position ('B')
  - Green: Solution path
  - Yellow: Explored cells (not part of solution)
  - White: Unvisited cells
  - Cell coordinates displayed on each cell
  - For informed search algorithms (Dijkstra, GBFS, A*), distance metrics are displayed on cells

- **Animation**: With `-animate` flag, generates `animation.png` (APNG format) showing the search process step-by-step:
  - Each frame stored temporarily in `tmp/` directory
  - Final animation combines all frames with the solution
  - Orange color indicates the current node being processed

- **Cell Size**: Images use 60x60 pixel cells (configurable via `cellSize` constant in `image.go`)

- **Grid Lines**: Gray grid lines separate cells for clarity

### Directory Structure

- `mazes/`: Contains maze text files for testing different scenarios
  - `maze.txt`: Basic maze
  - `maze2.txt`: Alternative maze
  - `maze-100-steps.txt`: Larger maze with longer solution path
- `tmp/`: Temporary storage for animation frames (auto-created and emptied on startup)

## Module Organization

This is a Go workspace project. When adding new modules or working with dependencies, use the workspace-aware commands and ensure `go.work` is updated if needed.

### Dependencies

The project uses the following external packages:
- `github.com/StephaneBunel/bresenham`: For drawing grid lines
- `github.com/kettek/apng`: For creating animated PNG files
- `golang.org/x/image/font`: For rendering text on images
