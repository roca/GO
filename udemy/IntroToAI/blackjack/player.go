package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Hit            = "h"
	Stand          = "s"
	Quit           = "q"
	MinDealerStand = 17
)

type Player struct {
	Name   string
	Hand   []Card
	Score  int
	IsAI   bool
	IsBust bool
}

func NewPlayer(name string, isAI bool) Player {
	return Player{
		Name:   name,
		Hand:   []Card{},
		Score:  0,
		IsAI:   isAI,
		IsBust: false,
	}
}

// CalulateScore calculates the player's score based on their hand of cards.
func (p *Player) CalculateScore() int {
	// First, add up all the non-ace cards.
	nonAceScore := 0
	aces := 0

	// Go through each card
	for _, card := range p.Hand {
		if card.Value == "A" {
			// Count the number of aces
			aces++
		} else {
			// Add the value of non-ace cards
			nonAceScore += card.Score
		}
	}

	// Handle aces if any
	acrScore := 0
	for range aces {
		// If adding 11 keeps the score <= 21, use 11, otherwise use 1
		if nonAceScore+acrScore+11 <= 21 {
			acrScore += 11
		} else {
			acrScore += 1
		}
	}

	return nonAceScore + acrScore
}

// AddCard adds a card to the player's hand and updates their score.
func (p *Player) AddCard(card Card, cardCounter *CardCounter) {
	// Add the card to the player's hand
	p.Hand = append(p.Hand, card)

	// Update the player's score
	p.Score = p.CalculateScore()

	// If we are keeping track of cards, update the card cardCounter
	if cardCounter != nil {
		cardCounter.TrackCard(card)
	}
}

// DisplayHand displays the player's hand and score.
func (p *Player) DisplayHand(hideSecoindCard bool) {
	cards := []string{}

	// Go through each card in the player's hand
	for i, card := range p.Hand {
		if hideSecoindCard && i > 0 {
			// If we're hidding the second card, show ?? insted
			cards = append(cards, "??")
		} else {
			cards = append(cards, card.String())
		}
	}

	// Print the player's name, hand and all their cards
	fmt.Printf("%s's hand: %s", p.Name, strings.Join(cards, " "))

	// Show their score (or ? if we're hiding cards)
	if hideSecoindCard {
		fmt.Printf(" (Score: ?)\n")
	} else {
		fmt.Printf(" (Score: %d)\n", p.Score)
	}
}

// handleHit
func (p *Player) handleHit(deck *Deck, cardCounter *CardCounter) bool {
	// Draw a card and add it to the player's hand
	card := deck.Draw()
	p.AddCard(card, cardCounter)

	// Show the card that they got
	fmt.Printf("%s drew: %s\n", p.Name, card.String())
	p.DisplayHand(false)

	// Check to see if they went over 21
	if p.Score > 21 {
		fmt.Printf("%s busts with a score over 21!\n", p.Name)
		p.IsBust = true
		return true
	}

	// If it's the AI's turn, add a small delay to make it easier to follow
	if p.IsAI {
		time.Sleep(1 * time.Second)
	}

	return false
}

// PlayTurn handles the player's turn, allowing them to hit or stand.
func (p *Player) PlayTurn(deck *Deck, cardCounter *CardCounter, dealerUpCard Card) {
	if p.IsAI {
		p.playAITurn(deck, cardCounter, dealerUpCard)
	} else {
		// If it's a human, let them choose what to do
		p.playHumanTurn(deck, cardCounter)
	}
}

func (p *Player) playHumanTurn(deck *Deck, cardCounter *CardCounter) {
	fmt.Printf("\n--- %s's Turn ---\n", p.Name)

	// Keep asking them what they want to do? (h)it, (s)tand or (q)uit
	for {
		fmt.Printf("What would you like to do? (h)it, (s)tand or (q)uit: ")
		var choice string
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)

		switch choice {
		case Quit:
			// They want to quit the game
			fmt.Println("Thanks for playing! Goodbye!")
			return
		case Hit:
			// Player wants to hit
			if p.handleHit(deck, cardCounter) {
				return
			}
		case Stand:
			// They're happy with their cards
			fmt.Printf("%s chose to stand.\n", p.Name)
			return
		default:
			fmt.Println("Invalid choice. Please enter 'h' to hit, 's' to stand, or 'q' to quit.")
		}
	}
}

func (p *Player) playAITurn(deck *Deck, cardCounter *CardCounter, dealerUpCard Card) {
	fmt.Printf("\n--- %s's Turn ---\n", p.Name)

	// Keep going until the AI decides to stand or busts
	for !p.IsBust {
		//Ask the AI what it wants to do
		choice := AdvancedAIDecision(*p, dealerUpCard, cardCounter)

		if choice == Stand {
			fmt.Printf("%s chose to stand.\n", p.Name)
			break
		}

		if choice == Hit {
			if p.handleHit(deck, cardCounter) {
				break // AI busts, end their turn
			}
		}

		if len(p.Hand) > 10 {
			fmt.Printf("%s has too many cards, automatically standing.\n", p.Name)
			break
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	reader.ReadString('\n')
}
