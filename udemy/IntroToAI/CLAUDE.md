# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains code for a Udemy course: [Introduction to AI and Machine Learning with Go](https://www.udemy.com/course/introduction-to-ai-and-machine-learning-with-go-golang)

The repository uses Go workspaces (`go.work`) with two modules:
- `ai-search/` - Pathfinding algorithms for maze solving (complete)
- `vacuum-1/` - Room cleaning robot simulator (work in progress)

**Requirements**: Go 1.25.6 or later

## Commands

### AI Search

```bash
cd ai-search
go run . -file mazes/maze.txt -search dfs
go run . -file mazes/maze-100-steps.txt -search astar -animate
go run . -file flooded-mazes/maze-flooded.txt -search astar
go build  # produces ./ai-search binary
```

Flags: `-file <path>`, `-search <algorithm>`, `-debug`, `-animate`

Algorithms: `dfs`, `bfs`, `dijkstra`, `gbfs`, `astar`

### Vacuum-1

```bash
cd vacuum-1
go run . -file empty.json -algorithm random -animate
go run . -file room.json -algorithm random -animate
```

Flags: `-file <path>`, `-algorithm <name>`, `-animate <bool>`

### Go Workspace

```bash
go work sync
go work use ./new-module
go work edit -dropuse ./module
```

### Testing

No test files exist in the codebase currently.

## Architecture: ai-search

### Core Types (main.go)

- **Maze**: Main game state - grid of Walls, Start/Goal Points, Solution, Explored list, search config
- **Point**: Row/Col coordinate with `Water` boolean (for flooded cells)
- **Wall**: Grid cell with `State` (Point) and `wall` boolean (`true` = impassable)
- **Node**: Search tree node with State, Parent pointer, Action, cost fields (`CostToGoal` for g(n), `EstimatedCostToGoal` for f(n)). Has `ManhattanDistance(goal Point) int` method.
- **Solution**: Result with Actions (direction strings) and Cells (path Points)

Search type constants: `DFS=0, BFS=1, GBFS=2, ASTAR=3, DIJKSTRA=4`

### Search Algorithm Pattern

Each algorithm follows the same structure (see `dfs.go` as reference):

1. **Struct** with `Frontier` (slice or priority queue) and `Game *Maze`
2. **Methods**: `Add(node)`, `Remove() (*Node, error)`, `ContainsState(node) bool`, `Empty() bool`, `Neighbors(node) []*Node`, `Solve()`
3. **Entry function**: `solve<Name>(m *Maze)` called from main.go dispatch

**Search logic**: Initialize frontier with start node → loop: remove node, check if goal, add unexplored neighbors → backtrack through parent pointers to build solution.

**IMPORTANT**: In `Neighbors()`, traversable cells have `wall == false`. Check `!wall`, not `wall`.

### Algorithm Differences

| Algorithm | Frontier | Ordering | Randomized | Optimal |
|-----------|----------|----------|------------|---------|
| DFS | Slice (LIFO stack) | Removes from end | Yes | No |
| BFS | Slice (FIFO queue) | Removes from front | No | Yes (shortest path) |
| Dijkstra | PriorityQueueDijkstra | g(n) - cost from start | No | Yes |
| GBFS | PriorityQueueGBFS | h(n) - Manhattan distance to goal | No | No |
| A* | PriorityQueueAstar | f(n) = g(n) + h(n), h = Euclidean distance | No | Yes |

DFS randomizes neighbor order, so multiple runs may produce different paths. All other algorithms use consistent ordering.

### Water/Flooded Cells

- Marked with `w` in maze files (see `flooded-mazes/`)
- `Point.Water` boolean set during maze loading
- A* adds `FloodedCost = 100` penalty in `Add()` when `node.State.Water` is true
- Water state propagated from maze grid to neighbor nodes in A*'s `Neighbors()`
- Rendered in blue in generated images

### Priority Queues

Three separate files implementing Go's `heap.Interface` (`Len`, `Less`, `Swap`, `Push`, `Pop`):
- `priority-queue-dijkstra.go`: Orders by `CostToGoal` (g(n))
- `priority-queue-gbfs.go`: Orders by `CostToGoal` (contains heuristic value h(n))
- `priority-queue-astar.go`: Orders by `EstimatedCostToGoal` (f(n))

### Image Generation (image.go)

- `OutputImage()`: Static `image.png` with 60x60 pixel cells
- `OutPutAnimatedImage()`: APNG `animation.png` from frames in `tmp/`
- Colors: Black=walls, Dark Green=start, Red=goal, Green=solution, Yellow=explored, Blue=water, White=unvisited, Orange=current node
- Informed search algorithms (Dijkstra, GBFS, A*) display distance metrics on cells

### Helper Functions (helpers.go)

- `inExplored()`: Check if Point was visited
- `emptyTmp()`: Clear animation frames from tmp/
- `abs()`: Integer absolute value
- `euclideanDist()`: Euclidean distance between Points (A* heuristic)

### Implementing a New Search Algorithm

1. **Create priority queue** (if needed): `priority-queue-<name>.go` implementing `heap.Interface`
2. **Create algorithm file** `<name>.go`:
   - Algorithm struct with `Frontier` and `Game *Maze`
   - Implement: `Add`, `Remove`, `ContainsState`, `Empty`, `Neighbors`, `Solve`
   - Entry function: `solve<Name>(m *Maze)`
3. **Update `main.go`**:
   - Add constant to search type enum (lines 14-19)
   - Add case to switch statement (lines 97-116)
4. **Update this CLAUDE.md**

### Maze File Format

Files in `ai-search/mazes/` and `ai-search/flooded-mazes/`:
- `A` = start, `B` = goal, `#` = wall, ` ` (space) = open path, `w` = water

## Architecture: vacuum-1

### Current Status

- **Complete**: Data structures, CLI parsing, room config loading (`LoadRoomConfig()`), room initialization (`NewRoom()`), robot constructor (`NewRobot()`), display system (`Display()`), furniture placement, main workflow, A* pathfinding in `astar.go`, summary display (`displaySummary()`), random walk core loop with movement, stuck detection, and A* fallback
- **Implemented helpers** (all in `random.go`): `bresenhamLine(x0, y0, x1, y1) []Point` (Bresenham's line algorithm), `moveAtAngleUntilObstacle(room, robot, dx, dy) int` (moves robot along a vector until hitting obstacle), `Abs(x) int` (integer absolute value), `findNearestDirtyCell(room, position) Point` (scans grid for nearest cell using Manhattan distance — note: currently doesn't filter by cell dirty status)
- **TODO**: Additional cleaning algorithms beyond random walk, cat obstacle behavior, loading robot dock position from JSON config

### Known Bugs

- `Clean()` in robot.go:33 increments `CleanableCellCount` instead of `CleanedCellCount` — room starts with correct count from `NewRoom()` but cleaning corrupts it
- `Display()` in world.go handles cell type `"clean"` but `Clean()` sets type to `"cleaned"` — cleaned cells won't render with any character
- Typo `"furnniture"` in robot.go:51 (should be `"furniture"`) — obstacle recording never matches furniture cells
- `CleanRoomRandomWalk()` in random.go:25 loop condition is `CleanableCellCount < CleanedCellCount` (inverted — should be `CleanedCellCount < CleanableCellCount`)
- `Display()` switch in world.go has no case for `"bike"` type — bike furniture cells render as blank
- `findNearestDirtyCell()` in random.go doesn't check if cells are actually dirty — scans all non-wall interior cells

### Constants (world.go)

- `cellSize = 10`: Room dimensions (cm) are divided by this to get grid dimensions (e.g., 300cm → 30 cells)
- `moveDelay = 50ms`: Animation delay between moves
- `catStopProbability = 0.1`, `catStopDuration = 5`: Cat obstacle behavior (not yet implemented)

### Key Types (world.go, robot.go)

- **Room**: Grid of Cells, dimensions, cleanable/cleaned counts, animate flag
- **Cell**: Type string ("clean"/"dirty"/"wall"/"furniture"/"bike"), Cleaned/Obstacle booleans
- **Robot**: Position, Path, Direction (float64), CleanRoom function pointer, ObstaclesEncountered map
- **RoomConfig**: JSON structure with Width, Height, Furniture array

**Grid convention**: `grid[x][y]` where x is column and y is row (column-major). Display iterates rows (j) then columns (i).

**Shared state**: The `directions` variable (N/E/S/W offsets) is defined in `robot.go` and used by both robot movement and `astar.go` pathfinding.

### Display Characters

🔴 Robot, 🟦 Wall, 🪑 Furniture, 🧽 Clean, 🟫 Dirty, 🟢 Path, 🐱 Cat

### Room Configuration

JSON files in `vacuum-1/`. Use `empty.json` as template.

The JSON `robot` field (dock position, start direction) exists in config files but is not loaded by `RoomConfig` struct.

### Pathfinding (astar.go)

A* implementation for room navigation, separate from the ai-search module's A*. Uses a `PriorityQueue` of `PGItem` structs with `container/heap`. The `heuristic()` function uses Manhattan distance. Complete implementation: maintains open/closed sets, explores 4-directional neighbors via `directions` (defined in robot.go), tracks g/f scores, and reconstructs path via `cameFrom` map. Called as `Astar(room, start, goal) []Point`.

### Cleaning Algorithms

Assigned to robots via function pointers. Currently only `CleanRoomRandomWalk` in `random.go`. The algorithm: generates a random angle, computes a direction vector, uses `moveAtAngleUntilObstacle()` to move along that vector cleaning cells, and tracks a stuck counter — when stuck too many times consecutively, falls back to A* pathfinding toward the nearest dirty cell. All helper functions (`bresenhamLine()`, `moveAtAngleUntilObstacle()`, `Abs()`, `findNearestDirtyCell()`) are in `random.go`.

### Execution Flow

1. Parse flags → 2. `NewRoom(configFile, animate)` → 3. `NewRobot(1, 1)` → 4. Assign algorithm → 5. `robot.CleanRoom(room, robot)`

## Dependencies

- `github.com/StephaneBunel/bresenham`: Grid line drawing
- `github.com/kettek/apng`: Animated PNG generation
- `golang.org/x/image/font`: Text rendering on images

## Git Notes

- Git LFS tracks `.png` and `.psd` files (see `.gitattributes`)
- `.gitignore` excludes `tmp/` and `ai-search/*.png` (generated output)
- Main branch: `main`, active development on `staging`
