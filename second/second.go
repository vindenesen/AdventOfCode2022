package second

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Round struct {
	Opponent *Hand
	Me       *Hand
}

type Hand struct {
	Shape int
}

func (h *Hand) String() string {
	if h.Shape == Rock {
		return "Rock"
	} else if h.Shape == Paper {
		return "Paper"
	} else if h.Shape == Scissor {
		return "Scissor"
	}
	return ""
}

// Beats
// Returns 1 = Win, 0 = Loose, -1 = Tie
// /*
func (h *Hand) Beats(opponent *Hand) int {
	if h.Shape == Rock {
		if opponent.Shape == Paper {
			fmt.Println("    Rock looses over paper")
			return Loose
		}
		if opponent.Shape == Scissor {
			fmt.Println("    Rock wins over scissor")
			return Win
		}
	} else if h.Shape == Scissor {
		if opponent.Shape == Paper {
			fmt.Println("    Scissor wins over paper")
			return Win
		}
		if opponent.Shape == Rock {
			fmt.Println("    Scissor looses over rock")
			return Loose
		}
	} else if h.Shape == Paper {
		if opponent.Shape == Scissor {
			fmt.Println("    Paper looses over scissor ")
			return Loose
		}
		if opponent.Shape == Rock {
			fmt.Println("    Paper wins over rock")
			return Win
		}
	}

	fmt.Println("    Tie")
	return Tie
}

const (
	inputFile = "second/input"

	Win   = 1
	Loose = -1
	Tie   = 0

	Rock    = 1
	Paper   = 2
	Scissor = 3

	ScoreWinner = 6
	ScoreTie    = 3
	ScoreLooser = 0
)

// GetScore
// Returns (my score, opponents score)
// /*
func (r *Round) GetScore() (int, int) {
	myScore := ScoreLooser
	opponentScore := ScoreLooser

	if r.Me.Beats(r.Opponent) == Win {
		myScore = ScoreWinner + r.Me.Shape
		opponentScore = r.Opponent.Shape
		fmt.Println("   I won, with score", myScore)
	} else if r.Me.Beats(r.Opponent) == Loose {
		opponentScore = ScoreWinner + r.Opponent.Shape
		myScore = r.Me.Shape
		fmt.Println("   Opponent won, with score", opponentScore)
	} else {
		myScore = ScoreTie + r.Me.Shape
		opponentScore = ScoreTie + r.Opponent.Shape
		fmt.Println("   Tie, with score", myScore, opponentScore)
	}
	return myScore, opponentScore
}

func (r *Round) String() string {
	return r.Opponent.String() + " " + r.Me.String()
}

func GetRounds() []*Round {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	rounds := make([]*Round, 0)

	for _, l := range strings.Split(string(content), "\n") {
		round := Round{}
		if l != "" {
			letters := strings.Split(l, " ")
			round.Opponent = parseOpponentHand(letters[0])
			round.Me = parseMyHand(letters[1])
			rounds = append(rounds, &round)
		}
	}

	return rounds
}

func parseOpponentHand(char string) *Hand {
	if char == "A" {
		return &Hand{Shape: Rock}
	} else if char == "B" {
		return &Hand{Shape: Paper}
	} else if char == "C" {
		return &Hand{Shape: Scissor}
	}
	return nil
}

func parseMyHand(char string) *Hand {
	if char == "X" {
		return &Hand{Shape: Rock}
	} else if char == "Y" {
		return &Hand{Shape: Paper}
	} else if char == "Z" {
		return &Hand{Shape: Scissor}
	}
	return nil
}

func CalculateScore(rounds []*Round) (int, int) {
	myScore := 0
	opponentScore := 0

	for _, r := range rounds {
		fmt.Println("*", r)
		me, opponent := r.GetScore()
		myScore += me
		opponentScore += opponent
		fmt.Printf("Current scores, me %d, opponent %d\n", myScore, opponentScore)
	}

	return myScore, opponentScore
}

func printPartOne() {
	fmt.Println("--- Part one ---")
	rounds := GetRounds()
	fmt.Println("Rounds", len(rounds))
	myScore, opponentScore := CalculateScore(rounds)
	fmt.Printf("My score will be %d (opponent %d)\n\n", myScore, opponentScore)
}

func PrintAnswer() {
	fmt.Println("--- Day 2: Rock Paper Scissors ---")
	printPartOne()
}
