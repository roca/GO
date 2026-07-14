package main

import "testing"

// emptyBoard returns a pointer to a board with every cell set to empty,
// matching the state produced by the player constructors.
func emptyBoard() *Board {
	var b Board
	for i := range boardSize {
		for j := range boardSize {
			b[i][j] = empty
		}
	}
	return &b
}

func TestAbs(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
		{-1, 1},
	}
	for _, tt := range tests {
		if got := abs(tt.in); got != tt.want {
			t.Errorf("abs(%d) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestCheckWinCondition(t *testing.T) {
	t.Run("no ships remaining is a win", func(t *testing.T) {
		b := emptyBoard()
		if !checkWinCondition(b) {
			t.Error("checkWinCondition = false on a board with no ships, want true")
		}
	})

	t.Run("hits and misses still count as a win", func(t *testing.T) {
		b := emptyBoard()
		b[3][3] = hit
		b[4][4] = miss
		if !checkWinCondition(b) {
			t.Error("checkWinCondition = false with only hits/misses, want true")
		}
	})

	t.Run("a remaining ship cell is not a win", func(t *testing.T) {
		b := emptyBoard()
		b[0][0] = ship
		if checkWinCondition(b) {
			t.Error("checkWinCondition = true with a ship cell present, want false")
		}
	})
}

func TestIsShipSunkHorizontal(t *testing.T) {
	ships := []Ship{{
		ShipName:      "Cruiser",
		StartPosition: Position{row: 2, col: 3},
		EndPosition:   Position{row: 2, col: 5},
	}}

	t.Run("fully hit ship is sunk", func(t *testing.T) {
		b := emptyBoard()
		b[2][3], b[2][4], b[2][5] = hit, hit, hit

		sunk, name := isShipSunk(b, 2, 4, ships)
		if !sunk || name != "Cruiser" {
			t.Errorf("isShipSunk = (%v, %q), want (true, \"Cruiser\")", sunk, name)
		}
	})

	t.Run("partially hit ship is not sunk", func(t *testing.T) {
		b := emptyBoard()
		b[2][3], b[2][4] = hit, hit // [2][5] still afloat

		sunk, name := isShipSunk(b, 2, 4, ships)
		if sunk || name != "" {
			t.Errorf("isShipSunk = (%v, %q), want (false, \"\")", sunk, name)
		}
	})
}

func TestIsShipSunkVertical(t *testing.T) {
	ships := []Ship{{
		ShipName:      "Destroyer",
		StartPosition: Position{row: 0, col: 7},
		EndPosition:   Position{row: 1, col: 7},
	}}

	t.Run("fully hit ship is sunk", func(t *testing.T) {
		b := emptyBoard()
		b[0][7], b[1][7] = hit, hit

		sunk, name := isShipSunk(b, 1, 7, ships)
		if !sunk || name != "Destroyer" {
			t.Errorf("isShipSunk = (%v, %q), want (true, \"Destroyer\")", sunk, name)
		}
	})

	t.Run("partially hit ship is not sunk", func(t *testing.T) {
		b := emptyBoard()
		b[0][7] = hit // [1][7] still afloat

		sunk, _ := isShipSunk(b, 0, 7, ships)
		if sunk {
			t.Error("isShipSunk = true for a half-hit vertical ship, want false")
		}
	})
}

func TestParseCoord(t *testing.T) {
	t.Run("valid coordinates", func(t *testing.T) {
		tests := []struct {
			in   string
			want Position
		}{
			{"A0", Position{row: 0, col: 0}},
			{"J9", Position{row: 9, col: 9}},
			{"C5", Position{row: 5, col: 2}},
			{"A9", Position{row: 9, col: 0}},
		}
		for _, tt := range tests {
			got, err := parseCoord(tt.in)
			if err != nil {
				t.Errorf("parseCoord(%q) returned error: %v", tt.in, err)
				continue
			}
			if got != tt.want {
				t.Errorf("parseCoord(%q) = %+v, want %+v", tt.in, got, tt.want)
			}
		}
	})

	t.Run("invalid coordinates return an error", func(t *testing.T) {
		bad := []string{
			"",    // empty
			"A",   // too short
			"00",  // non-letter column
			"K0",  // column past J
			"AA",  // non-numeric row
			"A10", // row past 9
			"A-1", // negative row
			"@5",  // column before A
		}
		for _, in := range bad {
			if _, err := parseCoord(in); err == nil {
				t.Errorf("parseCoord(%q) = nil error, want an error", in)
			}
		}
	})
}

func TestIsShipSunkNoShipAtLocation(t *testing.T) {
	ships := []Ship{{
		ShipName:      "Cruiser",
		StartPosition: Position{row: 2, col: 3},
		EndPosition:   Position{row: 2, col: 5},
	}}
	b := emptyBoard()
	b[9][9] = hit

	sunk, name := isShipSunk(b, 9, 9, ships)
	if sunk || name != "" {
		t.Errorf("isShipSunk = (%v, %q) for a cell with no ship, want (false, \"\")", sunk, name)
	}
}
