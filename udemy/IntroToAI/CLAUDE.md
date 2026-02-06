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
- `gbfs` - Greedy Best-First Search (planned)
- `astar` - A* Search (planned)

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
  - `CostToGoal`: Integer cost (used by Dijkstra and A*)

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
   - `Frontier`: Slice of `*Node` representing the search frontier (or `PriorityQueueDijkstra` for Dijkstra)
   - `Game`: Pointer to the `Maze` being solved

2. **Required Methods**:
   - `Add(node *Node)`: Add a node to the frontier
   - `Remove() (*Node, error)`: Remove and return a node from the frontier (LIFO for DFS, FIFO for BFS, priority for Dijkstra)
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
   - Dijkstra (`dijkstra.go`): Does NOT randomize - uses priority queue based on `CostToGoal`.

6. **Algorithm Differences**:
   - DFS uses LIFO (Last In, First Out) - removes from end of frontier like a stack
   - BFS uses FIFO (First In, First Out) - removes from beginning of frontier like a queue
   - Dijkstra uses a priority queue - removes node with lowest cost
   - BFS guarantees the shortest path but may explore more nodes
   - DFS may find a solution faster but doesn't guarantee optimality
   - Dijkstra guarantees the shortest path and is optimal for weighted graphs

7. **Helper Functions**: `helpers.go` contains utility functions like `inExplored()` for checking if a Point has been visited.

8. **Main Entry Point**: `main.go` handles file loading, validation, and dispatches to the appropriate solve function based on the `-search` flag

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
  - For Dijkstra, Manhattan distance from start is displayed

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
