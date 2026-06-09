package main

import "testing"

func TestInitializeHeatMap(t *testing.T) {
	p := &AIPlayer{}
	p.initializeHeatMap()

	// A corner cell: (0+0)%2==0 gets the checkerboard bonus; its center
	// distance is boardSize (10), well beyond maxCenterDistance, so no
	// center bonus.
	wantCorner := baseProbability + checkerboardBonus
	if got := p.heatMap[0][0]; got != wantCorner {
		t.Errorf("heatMap[0][0] = %d, want %d", got, wantCorner)
	}

	// Center cell (5,5): (5+5)%2==0 -> checkerboard, distance 0 -> center bonus.
	wantCenter := baseProbability + checkerboardBonus + centerProximityBonus
	if got := p.heatMap[5][5]; got != wantCenter {
		t.Errorf("heatMap[5][5] = %d, want %d", got, wantCenter)
	}

	// An odd-parity cell away from center: base only.
	if got := p.heatMap[0][1]; got != baseProbability {
		t.Errorf("heatMap[0][1] = %d, want %d", got, baseProbability)
	}
}

// newTestAIPlayer builds an AIPlayer with an initialized heat map and full
// potentialShips tracking, without placing ships or touching stdin.
func newTestAIPlayer() *AIPlayer {
	p := &AIPlayer{}
	p.initializeHeatMap()
	p.potentialShips = make([]shipTracker, len(shipTypes))
	for i, s := range shipTypes {
		p.potentialShips[i].size = s.size
	}
	return p
}

func TestApplyHuntModeBoostSingleHit(t *testing.T) {
	p := newTestAIPlayer()
	opp := emptyBoard()

	hitPos := Position{row: 5, col: 5}
	opp[5][5] = hit
	p.hits = []Position{hitPos}
	p.huntMode = true

	p.updateHeatMap(opp)

	neighbors := []Position{{4, 5}, {6, 5}, {5, 4}, {5, 6}}
	for _, n := range neighbors {
		if p.heatMap[n.row][n.col] < huntModeBoost {
			t.Errorf("neighbor %+v heat = %d, want >= huntModeBoost (%d)",
				n, p.heatMap[n.row][n.col], huntModeBoost)
		}
	}
}

func TestApplyHuntModeBoostHorizontalLine(t *testing.T) {
	p := newTestAIPlayer()
	opp := emptyBoard()

	// Two aligned hits on row 5: cols 4 and 5. Only the endpoints
	// (col 3 and col 6) should receive the hunt-mode boost.
	opp[5][4], opp[5][5] = hit, hit
	p.hits = []Position{{row: 5, col: 4}, {row: 5, col: 5}}
	p.huntMode = true

	p.updateHeatMap(opp)

	if p.heatMap[5][3] < huntModeBoost {
		t.Errorf("left endpoint [5][3] heat = %d, want >= huntModeBoost (%d)",
			p.heatMap[5][3], huntModeBoost)
	}
	if p.heatMap[5][6] < huntModeBoost {
		t.Errorf("right endpoint [5][6] heat = %d, want >= huntModeBoost (%d)",
			p.heatMap[5][6], huntModeBoost)
	}

	// A cell off the line (above col 4) should not get the line boost.
	if p.heatMap[4][4] >= huntModeBoost {
		t.Errorf("off-line cell [4][4] heat = %d, unexpectedly boosted", p.heatMap[4][4])
	}
}

func TestUpdateHeatMapSkipsTargetedCells(t *testing.T) {
	// Characterization: updateHeatMap skips targeted (hit/miss) cells with a
	// `continue`, so they are never reset to baseProbability and never receive
	// a ship-fit bonus. The cell retains exactly its initializeHeatMap value.
	ref := &AIPlayer{}
	ref.initializeHeatMap()
	wantUnchanged := ref.heatMap[0][0] // base + checkerboard for the corner

	p := newTestAIPlayer()
	opp := emptyBoard()
	opp[0][0] = miss

	p.updateHeatMap(opp)

	if p.heatMap[0][0] != wantUnchanged {
		t.Errorf("targeted cell [0][0] heat = %d, want %d (unchanged by ship-fit)",
			p.heatMap[0][0], wantUnchanged)
	}
}

func TestPlaceShipsInvariants(t *testing.T) {
	// PlaceShips uses the global RNG, so assert invariants over many runs
	// rather than exact coordinates.
	for run := range 100 {
		p := NewAIPlayer()
		p.PlaceShips()

		if len(p.ships) != len(shipTypes) {
			t.Fatalf("run %d: placed %d ships, want %d", run, len(p.ships), len(shipTypes))
		}

		// Each ship has the correct length and lies in bounds.
		for _, s := range p.ships {
			length := abs(s.EndPosition.row-s.StartPosition.row) +
				abs(s.EndPosition.col-s.StartPosition.col) + 1
			wantLen := shipLengthByName(s.ShipName)
			if length != wantLen {
				t.Errorf("run %d: ship %q length = %d, want %d", run, s.ShipName, length, wantLen)
			}
			for _, pos := range []Position{s.StartPosition, s.EndPosition} {
				if pos.row < 0 || pos.row >= boardSize || pos.col < 0 || pos.col >= boardSize {
					t.Errorf("run %d: ship %q endpoint %+v out of bounds", run, s.ShipName, pos)
				}
			}
		}

		// Total ship cells on the board match the sum of ship lengths (no overlap).
		wantCells := 0
		for _, s := range shipTypes {
			wantCells += s.size
		}
		got := 0
		for i := range boardSize {
			for j := range boardSize {
				if p.board[i][j] == ship {
					got++
				}
			}
		}
		if got != wantCells {
			t.Errorf("run %d: %d ship cells on board, want %d (overlap?)", run, got, wantCells)
		}
	}
}

func shipLengthByName(name string) int {
	for _, s := range shipTypes {
		if s.name == name {
			return s.size
		}
	}
	return -1
}
