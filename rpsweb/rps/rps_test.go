package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case ROCK:
			fmt.Println("Player chose rock")
		case PAPER:
			fmt.Println("Player chose paper")
		case SCISSORS:
			fmt.Println("Player chose scissors")
		}
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("Computer choice: %s\n", round.ComputerChoice)
		fmt.Printf("Round result: %s\n", round.RoundResult)
		fmt.Printf("Computer choice int: %d\n", round.ComputerChoiceInt)
		fmt.Printf("Computer score: %s\n", round.ComputerScore)
		fmt.Printf("Player score: %s\n", round.PlayerScore)
		fmt.Println("------------------------------------------------------")
	}
}
