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

Planned search algorithms (only DFS currently implemented):
- `dfs` - Depth-First Search (implemented)
- `bfs` - Breadth-First Search (planned)
- `gbfs` - Greedy Best-First Search (planned)
- `astar` - A* Search (planned)
- `dijkstra` - Dijkstra's Algorithm (planned)

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

- **Point**: Row/column coordinate (used for positions in the maze)

- **Wall**: Cell in the maze grid with:
  - `State`: Point representing position
  - `wall`: boolean - `true` means impassable wall, `false` means open path

- **Node**: Search tree node containing:
  - `State`: Point representing position
  - `Parent`: Pointer to parent node (for path reconstruction)
  - `Action`: String describing how we got here ("up", "down", "left", "right")

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
   - `Frontier`: Slice of `*Node` representing the search frontier
   - `Game`: Pointer to the `Maze` being solved

2. **Required Methods**:
   - `Add(node *Node)`: Add a node to the frontier
   - `Remove() (*Node, error)`: Remove and return a node from the frontier (LIFO for DFS, FIFO for BFS, etc.)
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

5. **Neighbor Ordering**: The DFS implementation randomizes neighbor order before adding them to the frontier. This means multiple runs may explore different paths even with the same maze.

6. **Helper Functions**: `helpers.go` contains utility functions like `inExplored()` for checking if a Point has been visited.

7. **Main Entry Point**: `main.go` handles file loading, validation, and dispatches to the appropriate solve function based on the `-search` flag

## Module Organization

This is a Go workspace project. When adding new modules or working with dependencies, use the workspace-aware commands and ensure `go.work` is updated if needed.
