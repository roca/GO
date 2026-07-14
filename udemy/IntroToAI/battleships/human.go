package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Ship struct {
	ShipName      string
	StartPosition Position
	EndPosition   Position
}

type HumanPlayer struct {
	board    Board
	ships    []Ship
	opponent *AIPlayer
}

func NewHumanPlayer() *HumanPlayer {
	return &HumanPlayer{board: *newBoard()}
}

func (p *HumanPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter target position (e.g., A0): ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(strings.ToUpper(input))

		pos, err := parseCoord(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		row, col := pos.row, pos.col

		if opponentBoard[row][col] == hit || opponentBoard[row][col] == miss {
			fmt.Printf("You have already targeted position %c%d. Please choose a different target.\n", 'A'+col, row)
			continue
		}

		// Determine if it's a hit or miss
		isHit := opponentBoard[row][col] == ship

		if isHit {
			opponentBoard[row][col] = hit
			fmt.Printf("Hit at %c%d!\n", 'A'+col, row)

			// Check if ship is sunk
			sunk, shipName := isShipSunk(opponentBoard, row, col, p.opponent.ships)
			if sunk {
				fmt.Printf("You sunk the opponent's %s!\n", shipName)
			}
		} else {
			opponentBoard[row][col] = miss
			fmt.Printf("Miss at %c%d.\n", 'A'+col, row)
		}

		time.Sleep(1 * time.Second)
		return Position{row: row, col: col}, true
	}
}

func (p *HumanPlayer) GetBoard() *Board {
	return &p.board
}

func (p *HumanPlayer) PlaceShips() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n== SHIP PLACEMENT ==")
	fmt.Println("Place your ships on the board.")
	fmt.Println("Format: A0 H (A0 starting position, H=horizontal or V=vertical)")
	fmt.Println("Positions are given as letter (A-J) for column and number (0-9) for row.")
	fmt.Println("Press Enter to continue...")
	reader.ReadString('\n')

	for _, shipType := range shipTypes {
		for {
			// Display current board state
			printBoards(&p.board, &Board{})

			fmt.Printf("\nPlace your %s (length %d): ", shipType.name, shipType.size)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToUpper(input))
			parts := strings.Fields(input)

			if len(parts) != 2 {
				fmt.Println("Invalid input format. Use format like 'A0 H'. Please try again.")
				time.Sleep(2 * time.Second)
				continue
			}

			posStr := parts[0]
			dir := parts[1]

			if dir != "H" && dir != "V" {
				fmt.Println("Invalid input format. Use format like 'A0 H'. Please try again.")
				time.Sleep(2 * time.Second)
				continue
			}

			pos, err := parseCoord(posStr)
			if err != nil {
				fmt.Println(err)
				time.Sleep(2 * time.Second)
				continue
			}
			row, col := pos.row, pos.col

			// Check if placement is valid
			valid := true
			positions := []Position{}

			for _, pos := range shipCells(Position{row: row, col: col}, shipType.size, dir == "H") {
				r, c := pos.row, pos.col

				// Check if ship would go off of board
				if r >= boardSize || c >= boardSize {
					valid = false
					fmt.Printf("Ship would go off of the board! (Attempted to place at position %c%d)\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}

				// Check if position overlaps with another ship
				if p.board[r][c] == ship {
					valid = false
					fmt.Printf("Ship overlaps with another ship at position %c%d! Please choose a different location.\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}

				positions = append(positions, pos)
			}

			if valid {
				// Place ship on board
				newShip := Ship{
					StartPosition: positions[0],
				}

				for _, pos := range positions {
					p.board[pos.row][pos.col] = ship
				}

				newShip.EndPosition = positions[len(positions)-1]
				newShip.ShipName = shipType.name
				p.ships = append(p.ships, newShip)
				break
			}
		}
	}

	//Show final placement
	printBoards(&p.board, &Board{})
	fmt.Println("All ships placed! Press Enter to start the game...")
	reader.ReadString('\n')
}
