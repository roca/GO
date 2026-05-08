package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	p := &HumanPlayer{}
	for i := range boardSize {
		for j := range boardSize {
			p.board[i][j] = empty
		}
	}
	return p
}

func (p *HumanPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter target position (e.g., A0): ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(strings.ToUpper(input))

		if len(input) < 2 {
			fmt.Println("Invalid input format. Please enter a letter (A-J) followed by a number (0-9).")
			continue
		}

		if input[0] < 'A' || input[0] > 'J' {
			fmt.Println("Invalid column. Please enter a letter between A and J.")
			continue
		}
		col := int(input[0] - 'A')

		rowStr := input[1:]
		row, err := strconv.Atoi(rowStr)
		if err != nil || row < 0 || row >= boardSize {
			fmt.Println("Invalid row. Please enter a number between 0 and 9.")
			continue
		}

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
			sunk, shipName := isShipSunk(opponentBoard, row, col, p, nil)
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

			pos := parts[0]
			dir := parts[1]

			if len(pos) < 2 || (dir != "H" && dir != "V") {
				fmt.Println("Invalid input format. Use format like 'A0 H'. Please try again.")
				time.Sleep(2 * time.Second)
				continue
			}

			// Extracting column (letter)
			if pos[0] < 'A' || pos[0] > 'J' {
				fmt.Println("Invalid column. Please enter a letter between A and J.")
				time.Sleep(2 * time.Second)
				continue
			}
			col := int(pos[0] - 'A')

			// Extractinmg row (number)
			rowStr := pos[1:]
			row, err := strconv.Atoi(rowStr)
			if err != nil || row < 0 || row >= boardSize {
				fmt.Println("Invalid row. Please enter a number between 0 and 9.")
				time.Sleep(2 * time.Second)
				continue
			}

			// Check if placement is valid
			valid := true
			positions := []Position{}

			for i := range shipType.size {
				var r, c int
				if dir == "H" {
					r, c = row, col+i // Horizontal placement increases column index
				} else {
					r, c = row+i, col // Vertical placement increases row index
				}

				// Check if ship would goo off of board
				if r >= boardSize || c >= boardSize {
					valid = false
					fmt.Printf("Ship wouild go off of the board! (Attemped to place at position %c%d)\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}

				//Check if position ovetrlaps with another ship
				if p.board[r][c] == ship {
					valid = false
					fmt.Printf("Ship overlaps with another ship at position %c%d! Please choose a different location.\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}

				positions = append(positions, Position{row: r, col: c})
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
