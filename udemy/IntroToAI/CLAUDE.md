# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains code for a Udemy course: [Introduction to AI and Machine Learning with Go](https://www.udemy.com/course/introduction-to-ai-and-machine-learning-with-go-golang)

The primary project is `ai-search`, which implements various pathfinding algorithms to solve maze problems.

## Project Structure

The repository uses Go workspaces (`go.work`) with the following modules:
- `ai-search/` - Main AI search algorithms project for maze solving
- `vacuum-1/` - Room cleaning robot simulator (work in progress)

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

Example with flooded/water cells:
```bash
go run main.go -file flooded-mazes/maze-flooded.txt -search astar
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
  - `Row`, `Col`: Integer coordinates
  - `Water`: Boolean flag indicating if this cell is flooded/water (affects cost in A*)

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

Maze files are stored in `ai-search/mazes/` and `ai-search/flooded-mazes/` as text files with:
- `A` - Start position
- `B` - Goal position
- `#` - Walls
- ` ` (space) - Open path
- `w` - Water/flooded cell (traversable but with high cost in A*)
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
   - **A***: Priority queue ordered by f(n) = g(n) + h(n), where g(n) is actual cost from start and h(n) is Euclidean distance to goal - removes node with lowest estimated total cost. **Special feature**: Water cells add a penalty cost (`FloodedCost = 100`) to discourage traversal through flooded areas
   - **Optimality**:
     - BFS guarantees shortest path but may explore more nodes (609 nodes for maze-100-steps.txt)
     - DFS may find a solution faster but doesn't guarantee optimality
     - Dijkstra guarantees shortest path and is optimal for weighted graphs (100 steps for maze-100-steps.txt)
     - GBFS is fast and explores fewer nodes (158 nodes for maze-100-steps.txt) but doesn't guarantee shortest path (108 steps)
     - A* guarantees shortest path (when using admissible heuristic) and is generally more efficient than Dijkstra. Can handle weighted terrain (water cells) by adding cost penalties

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

### Water/Flooded Cells Feature

The codebase supports weighted terrain through water/flooded cells:

- **Maze Representation**: Water cells are marked with `w` in maze files (see `flooded-mazes/` directory)
- **Cost Model**: In `astar.go`, the constant `FloodedCost = 100` adds a significant penalty to the estimated cost when traversing water cells
- **Implementation**:
  - The `Point` struct has a `Water` boolean field set during maze loading
  - In A*'s `Add()` method: if `node.State.Water` is true, add `FloodedCost` to `EstimatedCostToGoal`
  - In A*'s `Neighbors()` method: water state is propagated from the maze's Wall grid to the neighbor nodes
- **Visualization**: Water cells are rendered in blue in the generated images
- **Use Case**: Demonstrates how A* can handle weighted graphs where different terrain types have different traversal costs

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
  - Blue: Water/flooded cells (traversable with high cost)
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
- `flooded-mazes/`: Contains mazes with water cells (marked with 'w')
  - `maze-flooded.txt`: Maze with flooded/water areas
- `tmp/`: Temporary storage for animation frames (auto-created and emptied on startup)

## Vacuum-1 Module

The `vacuum-1` module is a room cleaning robot simulator (work in progress). It simulates autonomous robot vacuum cleaners navigating and cleaning rooms with obstacles.

**Current Status**: The vacuum-1 module has core infrastructure in place (room loading, grid creation, robot initialization, display system). The random walk cleaning algorithm exists as a stub in `random.go` and needs full implementation. Next steps include implementing movement logic, collision detection, and the actual random walk algorithm.

### Running Vacuum-1

```bash
cd vacuum-1
go run main.go -file empty.json -algorithm random -animate
```

Available command-line flags:
- `-file <path>`: Path to the room configuration file (default: "empty.json")
- `-algorithm <name>`: Cleaning algorithm to use (default: "random")
- `-animate <bool>`: Enable/disable animation while cleaning (default: true)

### Vacuum-1 Data Structures

All defined in `vacuum-1/world.go`:

- **Room**: Main simulation state containing:
  - `Grid`: 2D array of Cell structs representing the room
  - `Width`, `Height`: Room dimensions
  - `CleanableCellCount`: Total number of cells that can be cleaned
  - `CleanedCellCount`: Number of cells that have been cleaned
  - `Animate`: Boolean flag for animation

- **Cell**: Individual grid cell with:
  - `Type`: String ("clean", "dirty", "wall", "furniture", "bike")
  - `Cleaned`: Boolean indicating if cell has been cleaned
  - `Obstacle`: Boolean indicating if cell contains an obstacle
  - `ObstacleName`: String name of the obstacle if present

- **Point**: 2D coordinate with `X`, `Y` integer fields

- **Furniture**: Obstacle configuration with:
  - `X`, `Y`: Position coordinates
  - `Width`, `Height`: Dimensions
  - `Name`: Human-readable name (e.g., "Couch", "Coffee Table", "Bicycle")
  - `Type`: String type identifier

- **RoomConfig**: JSON configuration structure with:
  - `Width`, `Height`: Room dimensions
  - `Furniture`: Array of Furniture objects to place in the room

- **Robot**: Autonomous vacuum robot with:
  - `Position`: Point representing current location
  - `Path`: Slice of Points representing the path traveled
  - `CleanRoom`: Function pointer to the cleaning algorithm
  - `Direction`: Float64 representing the robot's orientation
  - `ObstaclesEncountered`: Map tracking encountered obstacles by name

- **NewRobot()**: Constructor function that creates a new robot
  - Takes `startX, startY` integers for initial position
  - Initializes `Position`, `Path`, and `ObstaclesEncountered` fields
  - Sets default `CleanRoom` function to `CleanRoomRandomWalk`
  - Returns `*Robot` pointer

### Display Characters

The simulator uses Unicode emoji for visualization:
- 🔴 (`charRobot`): Robot vacuum position
- 🟦 (`charWall`): Walls
- 🪑 (`charFurniture`): Furniture obstacles
- 🧽 (`charClean`): Cleaned cells
- 🟫 (`charDirty`): Dirty/uncleaned cells
- 🟢 (`carPath`): Robot's path
- 🐱 (`charCat`): Cat obstacle (with random stopping behavior)

### Room Configuration Files

Room configurations are stored as JSON files in `vacuum-1/`:
- `empty.json`: Empty room with just robot dock position
- `room.json`: Room with furniture obstacles (couch, coffee table, bicycle)

**Note**: The current `room.json` file has syntax errors (missing commas on lines 3, 17, 27, 36 and extra quote on line 38). Use `empty.json` as a template for valid JSON structure.

**Note**: The JSON configuration includes a `robot` field with dock position and start direction, but this is not yet used by the current implementation. The `RoomConfig` struct in `world.go` only loads `Width`, `Height`, and `Furniture` fields.

Configuration format:
```json
{
  "width": 300,
  "height": 300,
  "robot": {
    "diameter": 30,
    "dockX": 150,
    "dockY": 150,
    "startDirection": 0
  },
  "furniture": [
    {
      "id": 1,
      "x": 50,
      "y": 50,
      "width": 200,
      "height": 40,
      "name": "Couch",
      "type": "furniture"
    }
  ]
}
```

### Implementation Status

- ✅ Basic data structures defined (Room, Cell, Furniture, Point)
- ✅ Command-line argument parsing
- ✅ Room configuration loading (`LoadRoomConfig()` fully implemented)
- ✅ Room initialization (`NewRoom()` creates grid and walls)
- ✅ Robot struct defined with Position, Path, Direction, and CleanRoom function fields
- ✅ Display system (`Display()` renders room with emoji characters)
- ✅ Robot initialization (`NewRobot()` constructor implemented)
- ✅ Main workflow and algorithm dispatch (in `main.go`)
- ⏳ Furniture placement (TODO in world.go line 91)
- ⏳ Cleaning algorithms (random walk stub created in `random.go`, needs full implementation)
- ⏳ Animation/visualization system (display exists, animation loop needed)
- ⏳ Path planning and obstacle avoidance

### Constants

- `cellSize`: 10 pixels per cell
- `moveDelay`: 50 milliseconds between moves
- `catStopProbability`: 0.1 (10% chance cat stops moving each step)
- `catStopDuration`: 5 seconds when cat stops

### Room Display

The `Display()` method renders the current room state to the console:
- Clears the screen before each render
- Shows robot position, obstacles, and cell states using emoji characters
- Displays cleaning progress percentage at the bottom
- Optional `showPath` parameter to visualize the robot's traveled path

### Helper Functions

- `isInPath(Point, []Point) bool`: Check if a point exists in the robot's path (used for visualization)
- `displaySummary(room *Room, robot *Robot, moveCount int, cleaningTime time.Duration)`: Display final cleaning statistics including room size, coverage percentage, total moves, cleaning time, and efficiency (cells per move)

### Cleaning Algorithms

Cleaning algorithms are implemented in separate files and assigned to robots via function pointers:

- **random.go**: Contains `CleanRoomRandomWalk(*Room, *Robot)` function
  - Currently a stub implementation
  - Intended to implement random walk cleaning pattern
  - Calls `displaySummary()` when complete

### Main Execution Flow

The `main.go` file orchestrates the vacuum simulation:

1. Parse command-line flags (`-file`, `-algorithm`, `-animate`)
2. Create room using `NewRoom(configFile, animate)`
3. Create robot using `NewRobot(1, 1)` (starts at position 1,1)
4. Assign cleaning algorithm to robot based on `-algorithm` flag:
   - `"random"`: Uses `CleanRoomRandomWalk` function
5. Execute cleaning: `robot.CleanRoom(room, robot)`

## Module Organization

This is a Go workspace project. When adding new modules or working with dependencies, use the workspace-aware commands and ensure `go.work` is updated if needed.

### Requirements

- Go 1.25.6 or later (specified in go.work and module go.mod files)

### Working with Go Workspaces

Common workspace commands:
```bash
go work sync              # Sync workspace dependencies
go work use ./new-module  # Add a new module to the workspace
go work edit -dropuse ./module  # Remove a module from the workspace
```

### Dependencies

The project uses the following external packages:
- `github.com/StephaneBunel/bresenham`: For drawing grid lines
- `github.com/kettek/apng`: For creating animated PNG files
- `golang.org/x/image/font`: For rendering text on images
