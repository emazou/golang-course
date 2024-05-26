package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Rock. Beats scissors, loses to paper
	PAPER    = 1 // Paper. Beats rock, loses to scissors
	SCISSORS = 2 // Scissors. Beats paper, loses to rock
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Very nice! You won!",
	"Great job! You won!",
	"Awesome! You won!",
}

var loseMessages = []string{
	"Sorry, you lost!",
	"Better luck next time!",
	"Unlucky! You lost!",
}

var drawMessages = []string{
	"It's a draw!",
	"Nobody wins this round!",
	"Draw!",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "The computer chose rock"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "The computer chose paper"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "The computer chose scissors"
	}
	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "draw"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "win"
		message = winMessages[messageInt]
		PlayerScore++
	} else {
		roundResult = "lose"
		message = loseMessages[messageInt]
		ComputerScore++
	}
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
