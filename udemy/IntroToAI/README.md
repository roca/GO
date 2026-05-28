# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains code for a Udemy course: [Introduction to AI and Machine Learning with Go](https://www.udemy.com/course/introduction-to-ai-and-machine-learning-with-go-golang)

The repository uses Go workspaces (`go.work`) with five modules:
- `ai-search/` - Pathfinding algorithms for maze solving (complete)
- `vacuum-1/` - Room cleaning robot simulator (work in progress)
- `model-check/` - AI model fairness verification (functional — loads CSV, runs 3 model configs against fairness and risk properties)
- `battleships/` - Battleship game: human vs AI (in progress — ship placement, both players' turns, AI targeting/attack execution, and ship-sunk detection functional)
- `blackjack/` - Blackjack game with AI card counter (in progress — `Card`/`Deck`/`CardCounter` types and hi-lo card-counting helpers implemented; `main.go` round loop still placeholder comments)

**Requirements**: Go 1.25.6 or later (workspace declares 1.26.3 in `go.work`)

## Commands

### AI Search

```bash
cd ai-search
go run . -file mazes/maze.txt -search dfs
go run . -file mazes/maze-100-steps.txt -search astar -animate
go run . -file flooded-mazes/maze-flooded.txt -search astar
go build  # produces ./ai-search binary
```

Flags: `-file <path>`, `-search <algorithm>`, `-debug`, `-animate` (default: false)

Algorithms: `dfs`, `bfs`, `dijkstra`, `gbfs`, `astar`

### Vacuum-1

```bash
cd vacuum-1
go run . -file empty.json -algorithm random -animate
go run . -file room.json -algorithm random -animate
go run . -file room.json -algorithm slam -animate
go run . -file room.json -algorithm snake -animate
go run . -file room.json -algorithm snake -animate -cat
go run . -file house.json -algorithm snake -house
go run . -file house.json -algorithm snake -house -logic
```

Flags: `-file <path>`, `-algorithm <name>`, `-animate <bool>` (default: true), `-cat <bool>` (default: false), `-house <bool>` (default: false), `-logic <bool>` (default: false, use with `-house`)

Algorithms: `random`, `slam`, `spiral`, `snake` (default)

### Model Check

```bash
cd model-check
go run .
go build  # produces ./model-check binary
```

### Battleships

```bash
cd battleships
go run .
go build  # produces ./battleships binary
```

No flags — interactive console game (stdin/stdout).

### Blackjack

```bash
cd blackjack
go run .
go build  # produces ./blackjack binary
```

No flags — interactive console game. `main()` clears the screen, prints a welcome banner, builds a shuffled deck, and instantiates a `CardCounter`; the `for {}` round loop body is still placeholder comments.

### Build Verification

```bash
# From repo root, verify all modules compile:
cd ai-search && go build . && cd ../vacuum-1 && go build . && cd ../model-check && go build . && cd ../battleships && go build . && cd ../blackjack && go build .
```

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
- **Node**: Search tree node with State, Parent pointer, Action, cost fields (`CostToGoal` for g(n), `EstimatedCostToGoal` for f(n)), and `index` (used internally by priority queues). Has `ManhattanDistance(goal Point) int` method.
- **Solution**: Result with Actions (direction strings) and Cells (path Points)

Search type constants: `DFS=0, BFS=1, GBFS=2, ASTAR=3, DIJKSTRA=4`

### Search Algorithm Pattern

Each algorithm follows the same structure (see `dfs.go` as reference):

1. **Struct** with `Frontier` (slice or priority queue) and `Game *Maze`
2. **Methods**: `Add(node)`, `Remove() (*Node, error)`, `ContainsState(node) bool`, `Empty() bool`, `Neighbors(node) []*Node`, `Solve()`
3. **Entry function**: `solve<Name>(m *Maze)` called from main.go dispatch

**Search logic**: Initialize frontier with start node -> loop: remove node, check if goal, add unexplored neighbors -> backtrack through parent pointers to build solution.

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

### Maze File Format

Files in `ai-search/mazes/` and `ai-search/flooded-mazes/`:
- `A` = start, `B` = goal, `#` = wall, ` ` (space) = open path, `w` = water

### Implementing a New Search Algorithm

1. **Create priority queue** (if needed): `priority-queue-<name>.go` implementing `heap.Interface`
2. **Create algorithm file** `<name>.go`:
   - Algorithm struct with `Frontier` and `Game *Maze`
   - Implement: `Add`, `Remove`, `ContainsState`, `Empty`, `Neighbors`, `Solve`
   - Entry function: `solve<Name>(m *Maze)`
3. **Update `main.go`**:
   - Add constant to search type enum (see existing `DFS`, `BFS`, etc.)
   - Add case to algorithm switch statement in `main()`
4. **Update this CLAUDE.md**

## Architecture: vacuum-1

### Key Types

- **Room**: Grid of Cells, dimensions, cleanable/cleaned counts, animate flag, optional `Cat *Cat` pointer
- **Cell**: Type string ("clean"/"dirty"/"wall"/"furniture"/"bike"), Cleaned/Obstacle booleans, ObstacleName string
- **Furniture**: JSON-mapped struct with X, Y, Width, Height, Name, Type fields. Name is used for obstacle tracking via `ObstacleName`.
- **Robot**: Position, Path, Direction (float64), CleanRoom function pointer, ObstaclesEncountered map
- **House**: Contains `Rooms []*Room`. Created via `NewHouse(configFile, animate)`.
- **RoomConfig**: JSON structure with Width, Height, Furniture array. Single room JSON is one object; house JSON is an array of these.
- **Cat**: Position, Active bool, StopTimer, Path, DirectionX/DirectionY. Created via `NewCat(room)`, moved via `MoveCat(cat, room)`.
- **LogicalWorld**: Tracks Jack/Sarah/Johnny `PersonStatus` (IsHome, Room, DoorClosed), weekday flag, Objects map. Methods: `UpdateObjectFound()`, `UpdateDoorStatus()`, `DetermineCleaningPriority()`.
- **RobotWithLogic**: Embeds `*Robot`, adds `World *LogicalWorld`. Has `ScanHouseWithLogic(house)` method that identifies rooms by furniture, scans for objects, triggers deduction rules, and returns room-name-to-index mapping.

**Grid convention**: `grid[x][y]` where x is column and y is row (column-major). Display iterates rows (j) then columns (i).

**Shared state**: The `directions` variable (N/E/S/W offsets) defined in `robot.go` is used by both robot movement and `astar.go` pathfinding.

### Constants (world.go)

- `cellSize = 10`: Room dimensions (cm) are divided by this to get grid dimensions (e.g., 300cm -> 30 cells)
- `moveDelay = 50ms`: Animation delay between moves
- `catStopProbability = 0.1`, `catStopDuration = 5`: Used by `MoveCat()` in cat.go

### Pathfinding (astar.go)

A* implementation for room navigation, separate from the ai-search module's A*. Uses a `PriorityQueue` of `PQItem` structs with `container/heap`. The `heuristic()` function uses Manhattan distance (unlike ai-search's A* which uses Euclidean distance). Called as `Astar(room, start, goal) []Point`.

### Cleaning Algorithms

Assigned to robots via function pointers. `setAlgorithm()` in main.go maps names to functions: "random"->`CleanRoomRandomWalk`, "slam"->`CleanRoomSlam`, "spiral"->`CleanSpiralPattern`, default->`CleanRoomSnake`.

**Random Walk** (`random.go`): Random angle movement with stuck detection (5+ consecutive stuck -> A* fallback to nearest dirty cell). Every 20 moves, 30% chance of A* course correction. Final grid sweep for missed cells. Helper functions: `bresenhamLine()`, `moveAtAngleUntilObstacle()`, `findNearestDirtyCell()`, `abs()`.

**SLAM** (`slam.go`): Frontier-based exploration. Internal `robotMap` (0=unknown, 1=free, 2=obstacle, 3=cleaned). Picks closest frontier point, pathfinds with A*, expands frontier. Periodic re-scan every 10 moves. Early exit at 95% coverage. Final sweep via `cleanRemainingCells()`.

**Spiral** (`spiral.go`): Cleans from room center outward. `generateSpiralPattern()` creates expanding spiral (right->down->left->up). `findNearestCleanablePoint()` adjusts targets. `finalCleanup()` catches missed cells.

**Snake** (`snake.go`): Boustrophedon (zigzag) pattern. Default algorithm. `generateSnakingPattern()` alternates row direction. Only algorithm that calls `MoveCat()`. Reuses `finalCleanup()` from spiral.go.

### Execution Flow

1. Parse flags -> 2a. If `-house`: `NewHouse()` / 2b. Otherwise: `NewRoom()` wrapped in single-room `House` -> 3. Optionally `NewCat(room)` if `-cat` -> 4a. If `-logic`: `NewRobotWithLogic()` -> `setAlgorithm()` -> `ScanHouseWithLogic()` -> print logical state and priority -> wait for user input -> clean rooms in priority order / 4b. Otherwise: for each room -> `NewRobot()` -> `setAlgorithm()` -> `robot.CleanRoom(room, robot)`

### Cross-file Dependencies

- `snake.go` reuses `finalCleanup()` defined in `spiral.go` and calls `MoveCat()` from `cat.go`
- `directions` variable (N/E/S/W offsets) defined in `robot.go` is used by `astar.go`
- `cat.go` uses `abs()` from `random.go` and constants from `world.go`
- `robot.go`: `Clean()` calls `CheckAdjacentObstacles()` which uses `RecordObstacle()` to track furniture names

### Known Bugs

**Off-by-one errors in grid bounds**:
- `addNeighborsToFrontier()` in slam.go: `newX <= len(robotMap)` should be `<`
- `RecordObstacle()` in robot.go: `y <= room.Height` should be `y < room.Height`
- `finalCleanup()` in spiral.go: uses `room.Width-1`/`room.Height-1` as upper loop bounds, skipping last column and row
- `updateAllFrontiers()` in slam.go: same off-by-one with `Width-1`/`Height-1`

**Obstacle polarity**: `updateAllFrontiers()` in slam.go checks `room.Grid[x][y].Obstacle` when it should check `!room.Grid[x][y].Obstacle` to find free cells for the frontier.

**Out-of-bounds access**: `MoveCat()` in cat.go accesses `room.Grid[newX][newY]` outside the `room.IsValid()` guard block — will panic if position is invalid.

**Incomplete cell checks**: `findNearestDirtyCell()` in random.go finds nearest cell by distance but never checks `Cleaned` or `Obstacle` status — returns any cell, even if already clean or an obstacle.

**Typos**: `getCloesestFrontierPoint()` in slam.go (should be `getClosest`), `directiionX` in snake.go (double 'i').

**Logic bugs in spiral.go**: `findNearestCleanablePoint()` loop uses `||` where `&&` is needed. `generateSpiralPattern()` breaks on first out-of-bounds point, potentially missing valid points in non-square rooms.

**Logic branch missing `continue`**: In main.go, when `roomNameToIndex[roomName]` doesn't contain a priority room name, it prints "skipping" but doesn't `continue` — falls through to access `house.Rooms[0]` (zero value of missing map key).

### Incomplete Features (TODO)

- `-logic` flag: Scanning, priority determination, and room cleaning work end-to-end. Still needs: connect `RecordObstacle`/`CheckAdjacentObstacles` to `UpdateObjectFound` during cleaning, connect `UpdateDoorStatus` to door detection during cleaning.
- Cat integration: Only snake calls `MoveCat()`. Random, slam, and spiral need cat movement added. No algorithm calls `IsAdjacentToCat()` for avoidance.
- `Display()` switch in world.go has no case for `"bike"` cell type.
- JSON `robot` field (dock position, start direction) and furniture `id` field not loaded by `RoomConfig`/`Furniture` structs.

### Adding a New Cleaning Algorithm

1. Create `<name>.go` with a function matching signature `func(room *Room, robot *Robot)`
2. Add case to `setAlgorithm()` switch in `main.go`
3. If cat support needed, call `MoveCat()` during the cleaning loop
4. Update this CLAUDE.md

### Room Configuration

JSON files in `vacuum-1/`. Single-room files (`empty.json`, `room.json`) contain one `RoomConfig` object. Multi-room files (`house.json`) contain a JSON array of `RoomConfig` objects — use with `-house` flag.

## Architecture: model-check

Loan approval fairness verification module. Standard library only.

### Core Types (model.go)

- **LoanApprovalAI**: Weighted scoring model with factor weights (income, creditScore, loanAmount, debtRatio, employment) and `approvalThreshold`
- **Applicant**: Financial/demographic profile — income (thousands), creditScore (normalized 0-1), loanAmount (thousands), debtToIncome (0-1), yearsEmployed, protectedClass boolean

### Property System (property.go, risk-property.go)

- **Property**: Interface with `Check(model *LoanApprovalAI, applicants []Applicant) (bool, []Applicant)` and `Name() string`
- **FairnessProperty**: Implements `Property`, has `maxDisparity float64`. `Check()` counts approvals per group (protected vs non-protected), computes disparity, flags individually unfair decisions (denied despite strong profile: creditScore > 0.7, debtToIncome < 0.3, income > 60k), returns whether the model is fair and a list of unfair decisions.
- **RiskProperty**: Implements `Property`, has `axHighRiskApprovalRate float64` (note: field name missing `m` prefix — should be `maxHighRiskApprovalRate`). `Check()` identifies high-risk applicants (creditScore < 0.5 and debtToIncome > 0.5), computes their approval rate, returns whether it's within the allowed threshold plus the list of risky approvals.

### Implemented

- **`ApproveLoan`** (model.go): Weighted scoring — computes `loanToIncomeRatio`, then `score = income*w1 + creditScore*w2 - loanToIncomeRatio*w3 - debtToIncome*w4 + yearsEmployed*w5`, approves if `score > approvalThreshold`
- **`PrintModelParams`** (model.go): Displays model weights and threshold
- **`FairnessProperty.Check()`** (property.go): Full implementation — approval rate disparity check plus individual unfair-decision detection
- **`RiskProperty.Check()`** (risk-property.go): Full implementation — high-risk approval rate check
- **CSV loading** (`load-csv.go`): `LoadApplicantsFromCSV` with fuzzy column matching (header substring matching, case-insensitive) and auto-normalization (income/loanAmount to thousands, creditScore from 300-850 to 0-1, debtToIncome from percentage to ratio)
- **Parsing utilities** (`utils.go`): `parseFloat` (strips `$`, `,`, `%`) and `parseBool` (accepts true/yes/y/1/t/false/no/n/f/0)

### Verification (verification.go)

- **`VerifyModel`**: Takes model, property, and applicants. Calls `property.Check()`, prints pass/fail result with up to 3 counter-example details (income, credit score, debt ratio, protected class, decision).

### Known Bugs

- `risk-property.go`: Field name `axHighRiskApprovalRate` appears to be missing the `m` prefix (`maxHighRiskApprovalRate`)
- `risk-property.go`: Variable `isHoighRisk` is a typo (should be `isHighRisk`)
- `verification.go`: `Dept Ratio` should be `Debt Ratio` in format string
- `verification.go`: `theere` should be `there` in comment
- `main.go`: Comment says "Load applicant data from CSBV" — should be "CSV"

## Architecture: battleships

Classic Battleship game — human vs AI, interactive console. Standard library only. In progress: ship placement complete, both players' turns functional, AI targeting and attack execution implemented (hit/miss, hunt mode entry, sunk handling). `isShipSunk` is fully implemented; remaining gap is adjacency validation in human ship placement.

### Core Types

- **Board**: `[10][10]string` — 10x10 grid. Cells: `"."` (empty), `"O"` (ship), `"X"` (hit), `"~"` (miss)
- **Position**: `row`, `col` int pair (0-9)
- **Ship** (human.go): `ShipName`, `StartPosition`, `EndPosition`
- **HumanPlayer** (human.go): `board Board`, `ships []Ship`, `opponent *AIPlayer`. Constructor: `NewHumanPlayer()` initializes board with `"."`
- **AIPlayer** (ai.go): `board Board`, `heatMap [10][10]int`, `hits []Position`, `shipsSunk int`, `huntMode bool`, `potentialShips []struct{...}` with size/sunk/hits/shipPos fields, `ships []Ship`, `opponent *HumanPlayer`. Constructor: `NewAIPlayer()` initializes board, heat map, and potentialShips tracking from shipTypes.

### Helpers (helpers.go)

- `abs(x int) int`: Integer absolute value, used by heat map center-distance calculation
- `checkWinCondition(board *Board) bool`: Returns true when no `ship` cells remain on the board (all ships sunk)
- `isShipSunk(board *Board, row, col int, opponentShips []Ship) (bool, string)`: Locates the ship covering `(row, col)` by checking each opponent ship's `StartPosition`/`EndPosition` span (horizontal or vertical), counts `hit` cells along that span, and returns `(true, ShipName)` when the hit count equals the ship length, otherwise `(false, "")`. Called by both `HumanPlayer.TakeTurn` and `AIPlayer.TakeTurn` after a hit.

### Ship Registry (board.go)

`var shipTypes`: Carrier (5), Battleship (4), Cruiser (3), Submarine (3), Destroyer (2) — 5 ships, 17 total cells

### Constants (main.go)

`boardSize = 10`, display symbols (`empty="."`, `ship="O"`, `hit="X"`, `miss="~"`, `hiddenShip="."` — same glyph as empty so AI ships are indistinguishable from empty cells in the player's view), `headerRow = "  A B C D E F G H I J"`, `headerCol = "0123456789"`

### Game Loop (main.go)

Creates `HumanPlayer` and `AIPlayer`, links opponents, prints welcome/legend, calls `ai.PlaceShips()` then `human.PlaceShips()`, then enters alternating-turn `gameOver` loop. Before each AI turn, the loop prints the AI's `heatMap` as a 10×10 grid for debugging, then calls `ai.TakeTurn(human.GetBoard())`. The pause-for-Enter prompt lives inside `AIPlayer.TakeTurn` itself, not in the loop. Both turns call `checkWinCondition`. Final message on game end.

### AI Strategy Infrastructure

Two-phase hunt architecture with probability-based targeting:
- **Search phase**: `heatMap` initialized with checkerboard pattern + center bias via `initializeHeatMap()`. Constants: `baseProbability=1`, `checkerboardBonus=1`, `centerProximityBonus=2`, `maxCenterDistance=3`, `huntModeBoost=100`, `shipFitBonus=2`
- **Kill phase**: `huntMode` flag and `potentialShips` tracking defined; switches to targeted mode after hit
- **`updateHeatMap`**: Implemented — resets heat map, calculates base probabilities for untargeted cells, adds ship-fit bonuses (checks if unsunk ships can fit horizontally/vertically from each cell), calls `applyHuntModeBoosts()` when in hunt mode
- **`applyHuntModeBoosts`**: Implemented — detects hit pattern (single, horizontal line, or vertical line), boosts adjacent cells accordingly. Single hits boost all 4 neighbors; aligned hits boost only the endpoints of the line.
- **`TakeTurn`**: Implemented — calls `updateHeatMap`, finds highest probability cell(s), randomly selects among ties, falls back to random if no candidates. Resolves the attack: marks `hit`/`miss` on the opponent board, appends to `p.hits` and sets `huntMode = true` on a hit, calls `isShipSunk` and on a sunk ship increments `shipsSunk`, exits hunt mode, and clears `p.hits`. After printing the result, blocks on a "Press enter to continue..." prompt via `bufio.NewReader(os.Stdin)`. Returns the targeted `Position` and a `bool` indicating hit.

### AI Ship Placement (ai.go)

`PlaceShips()` is implemented — two-phase strategy: larger ships (size >= 4) placed near edges, smaller ships distributed randomly. Validates boundary, overlap, and adjacency (larger ships avoid diagonal/adjacent neighbors). Falls back to random valid placement after 100 failed attempts.

### Board Display (board.go)

`printBoards(playerBoard, opponentBoard *Board)`: Clears terminal, prints both boards side-by-side. Hides opponent's ships (shows `hiddenShip` instead of `ship`), reveals hits and misses. Player's own board shows all information.

### Known Bugs

- `ai.go`: Comment says `checkkerboard` (double 'k'), `likeley` (should be `likely`), `Sigificant` (should be `Significant`), `potenial` (should be `potential`), `hightest` (should be `highest`), `randowm` (should be `random`)
- `ai.go`: Variable `shipTyper` in range loop (should be `shipType`)
- `board.go`: Comment typos — `pacakage` (should be `package`), `opponewnt's` (should be `opponent's`), `shouild` (should be `should`)
- `human.go`: Comment typo `Extractinmg` (should be `Extracting`), `goo` (should be `go`), `wouild` (should be `would`), `Attemped` (should be `Attempted`), `ovetrlaps` (should be `overlaps`)

### Human Player (human.go)

`PlaceShips()` — prompts user with format "A0 H" (column letter + row number + direction), validates input format and direction (H/V), converts coordinates to Position (col from letter A-J, row from number 0-9), checks boundary and overlap, marks board cells, stores Ship with start/end positions, and displays final placement. Not yet implemented: adjacency checking (ships can be placed touching each other).

`TakeTurn(opponentBoard *Board) (Position, bool)` — prompts for target position (e.g., "A0"), validates input, checks for already-targeted cells, determines hit/miss, updates board, calls `isShipSunk` on hits. Returns the targeted position.

`GetBoard() *Board` — returns pointer to player's board.

### Not Yet Implemented

- Adjacency validation in human ship placement

## Architecture: blackjack

Console blackjack vs. an AI card counter. Partially implemented.

### Core Types

- **Card** (`card.go`): `Suit` (Unicode glyph constant — `Hearts`, `Diamonds`, `Clubs`, `Spades`), `Value` (`"A"`, `"2"`–`"10"`, `"J"`, `"Q"`, `"K"`), `Score` int. `String()` renders as `value+suit` (e.g., `A♠`). Ace is hard-coded to score 11 in `NewDeck` — no soft/hard-ace handling yet.
- **Deck** (`deck.go`): `[]Card`. `NewDeck()` builds a 52-card deck (suits × values, parallel `scores` slice). `Shuffle()` returns a Fisher-Yates-shuffled copy (does not mutate the receiver). `Draw()` is a pointer receiver that auto-reshuffles when empty — but **returns the top card without removing it**, so repeated calls yield the same card.
- **CardCounter** (`card-counter.go`): Tracks hi-lo card counting state. Fields: `SeenCards map[string]int` (per-value counts), `RunningCount int`, `TrueCount float64`, `DecksRemaining float64`. Constant `DeckSize = 52`.

### Card Counter API (`card-counter.go`)

- `NewCardCounter()`: Initializes counts (all values to 0), running/true counts to 0, decks remaining to 1.0.
- `Reset()`: Restores the counter to a fresh-deck state.
- `TrackCard(card Card)`: Increments per-value count, updates running count via hi-lo (low cards 2–6: `+1`; high cards 10/J/Q/K/A: `-1`; 7–9 unchanged), recomputes `DecksRemaining` (clamped to ≥ 0.1) and `TrueCount = RunningCount / DecksRemaining`.
- `ChanceOfBusting(playerScore int) float64`: Walks every card value, treats face cards as 10 and Ace as 11, counts unseen cards (4 per value minus seen, clamped to ≥ 0), tallies which would push the player over 21, returns the bust ratio. Returns `1.0` when `playerScore >= 21`, `0.5` when no unseen cards remain.
- `DealerChanceOfBusting(dealerUpCard Card) float64`: Looks up a base bust probability per dealer up-card, then adjusts by `TrueCount * 0.02`.

### Helpers (`helpers.go`)

- `clearScreen()`: On Windows uses `github.com/inancgumus/screen`; elsewhere writes the ANSI escape `\033[H\033[2J`.

### Game Loop (`main.go`)

`main()` clears the screen, prints a welcome banner, calls `NewDeck().Shuffle()`, and creates a `CardCounter` via `NewCardCounter()`. The `for {}` loop body is still placeholder comments — round play, shuffle-trigger, replay prompt, and exit condition are all unimplemented (loop is currently infinite).

### Known Bugs

- `Deck.Draw()` returns `(*d)[0]` but never slices it off, so the deck is never consumed and `Draw` always returns the same card.
- `NewDeck` assigns Ace `Score: 11` unconditionally; soft/hard-ace logic isn't implemented.
- `main` loop has no exit, so `go run .` will never return.
- `card-counter.go`: `pontsUntilBust` is a typo (should be `pointsUntilBust`); comments contain `caust` (should be `cause`), `somehowq` (should be `somehow`), `This  card` (double space).

## Dependencies

**ai-search** (external):
- `github.com/StephaneBunel/bresenham`: Grid line drawing
- `github.com/kettek/apng`: Animated PNG generation
- `golang.org/x/image/font`: Text rendering on images

**vacuum-1**: Standard library only

**model-check**: Standard library only

**battleships**: Standard library only

**blackjack** (external):
- `github.com/inancgumus/screen`: Terminal clearing on Windows (non-Windows uses ANSI escapes directly)

## Git Notes

- Git LFS tracks `.png` and `.psd` files (see `.gitattributes`)
- `.gitignore` excludes `tmp/`, `ai-search/*.png` (generated output), and compiled binaries (`ai-search/ai-search`, `vacuum-1/vacuum-1`)
- Main branch: `main`, active development on `staging`

**Note**: When adding or changing any algorithm (search or cleaning), update the relevant CLAUDE.md sections to keep documentation in sync.
