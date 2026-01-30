# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains code for a Udemy course on Introduction to AI and Machine Learning with Go. The primary project is `ai-search`, which implements various pathfinding algorithms to solve maze problems.

## Project Structure

The repository uses Go workspaces (`go.work`) with a single module:
- `ai-search/` - Main AI search algorithms project implementing DFS, BFS, GBFS, A*, and Dijkstra for maze solving

## Commands

### Running the AI Search Program

```bash
cd ai-search
go run main.go -file mazes/maze.txt -search dfs
```

Available search algorithms:
- `dfs` - Depth-First Search
- `bfs` - Breadth-First Search
- `gbfs` - Greedy Best-First Search
- `astar` - A* Search
- `dijkstra` - Dijkstra's Algorithm

### Building

```bash
cd ai-search
go build
```

## Architecture

### Maze Data Structure

The core data structures in `ai-search/main.go`:

- **Maze**: Represents the entire maze with dimensions, start/goal points, and wall grid
  - `Height`, `Width`: Maze dimensions
  - `Start`, `Goal`: Points marking the start ('A') and end ('B') positions
  - `Walls`: 2D array of Wall structs representing the maze grid

- **Point**: Simple row/column coordinate representation

- **Wall**: Represents a single cell in the maze with its position and whether it's a wall

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

### Search Algorithm Implementation

The program is structured to support multiple search algorithms (constants defined: DFS, BFS, GBFS, AStar, DIJKSTRA). The current implementation loads and validates mazes. Search algorithms will operate on the `Maze.Walls` grid, navigating from `Start` to `Goal` while avoiding walls (`wall = true`).

## Module Organization

This is a Go workspace project. When adding new modules or working with dependencies, use the workspace-aware commands and ensure `go.work` is updated if needed.
