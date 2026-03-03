# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go workspace containing multiple Fyne GUI toolkit projects, following the Udemy course [Fyne Beginner](https://www.udemy.com/course/fyne-beginner). Each subdirectory is an independent Go module managed via `go.work`.

## Build and Run Commands

The workspace uses Go 1.24.5 with Go workspaces (`go.work`).

```bash
# Run a specific project
cd project1 && go run .
cd project2 && go run .
cd project3 && go run .

# Build a specific project
cd project1 && go build -o project1 .

# Tidy dependencies for a specific module
cd project1 && go mod tidy

# Fyne packaging (project2 has FyneApp.toml configured)
cd project2 && fyne package
```

Fyne requires system-level graphics dependencies (OpenGL). On macOS these are available by default; on Linux install `libgl1-mesa-dev xorg-dev`.

## Architecture

- **go.work** — Workspace root linking `./myapp`, `./project1`, `./project2`, `./project3`
- **myapp/** — Fyne setup/scaffold module (no Go source, contains `Fyne Setup.app` bundle)
- **project1/** — Minimal Fyne app: single window with a label
- **project2/** — Interactive Fyne app: VBox layout with label, text entry, and button. Has `FyneApp.toml` metadata and `Icon.png` for packaging
- **project3/** — Kubeconfig Viewer: tabbed Fyne app that parses `~/.kube/config` (via `gopkg.in/yaml.v3`) and displays clusters, contexts, and users in scrollable card layouts with the current context highlighted

All projects use `fyne.io/fyne/v2` (project1/2: v2.6.x, project3: v2.7.x). Projects are standalone `main` packages with no shared library code between them.

## Fyne Patterns Used

- `app.New()` / `a.NewWindow()` for app lifecycle
- `widget.NewLabel`, `widget.NewEntry`, `widget.NewButton` for UI elements
- `container.NewVBox` for vertical layout composition
- `container.NewAppTabs` / `container.NewTabItem` for tabbed navigation (project3)
- `container.NewVScroll` for scrollable content areas (project3)
- `widget.NewCard` for grouped display sections (project3)
- `widget.NewForm` / `widget.NewFormItem` for key-value display (project3)
- `canvas.NewText` for styled text with custom color and size (project3)
- Callback-based event handling (button `func()` closures that update widget state)
