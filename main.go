package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// size of the board, always a square, ex: 4 means 4x4
const BOARD_SIZE int = 4

// store the information for a tile in a simple struct
// lmao i'm not using it right now, maybe if i need to store other info later?

// type tile struct {
// 	x, y, val int
// }

func main() {
	// randomize the seed because it isn't random by default
	seed := time.Now().Unix()
	rand.Seed(seed)

	fmt.Println("Welcome to GO 2048!")

	gameBoard := [BOARD_SIZE][BOARD_SIZE]int{}

	// game loop omg like python games course 😱
	for {
		if addTile(&gameBoard) == false {
			fmt.Printf("\nyou lost dumbass 🤡\n")
			return
		}

		printBoard(&gameBoard)

		input := getInput()
		move(input, &gameBoard)
	}
}

func printBoard(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// prints the board i dont know what else to say
	fmt.Printf("\n")
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getInput() string {
	// gets input from the user, either u, d, l, or r
	var userInput string

	for !isValid(&userInput) {
		fmt.Printf("Enter an input (u/d/l/r): ")
		fmt.Scanf("%s", &userInput)
	}

	return userInput
}

func isValid(s *string) bool {
	// takes in a string s and checks if it is a valid input, returns a bool
	// valid inputs start with the letter u, d, l, or r (case insensitive)
	validInputs := [4]string{"u", "d", "l", "r"}

	if len(*s) == 0 {
		fmt.Println("Please enter an input")
		return false
	}
	*s = strings.ToLower(*s)
	*s = (*s)[:1]
	for i := 0; i < len(validInputs); i++ {
		if *s == validInputs[i] {
			return true
		}
	}
	fmt.Println("Invalid input entered")
	return false
}

func addTile(b *[BOARD_SIZE][BOARD_SIZE]int) bool {
	// takes in the board and returns true if it adds a tile successfully
	// returns false if unsuccessful aka the board is full

	// create a slice of all of the empty tiles (tiles with a value of 0)
	emptyTiles := make([]*int, 0, 16)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (*b)[i][j] == 0 {
				emptyTiles = append(emptyTiles, &(*b)[i][j])
			}
		}
	}

	if len(emptyTiles) == 0 {
		return false
	}

	// pick a random tile from the empty tiles
	tile := emptyTiles[rand.Intn(len(emptyTiles))]

	// pick a value, mostly 2, 10% chance it is a 4
    val := 2
	percent := rand.Intn(10)
	if percent == 0 {
		val = 4
    }

	// assign the value to the tile
	*tile = val

	return true
}

// move code below to separate file eventually, or not idk
func move(input string, b *[BOARD_SIZE][BOARD_SIZE]int) {
	// how tf am i going to make the tiles move 😭
	switch input {
	case "u":
		fmt.Printf("TODO LMAO\n")
	case "d":
		fmt.Printf("TODO LMAO\n")
	case "l":
		moveLeft(b)
	case "r":
		moveRight(b)
	}
}

func moveLeft(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (*b)[i][j] != 0 {
				for pos := j; pos != 0; pos-- {
					currTile := &(*b)[i][pos]
					nextTile := &(*b)[i][pos-1]

					if *nextTile == 0 { // slide
						*nextTile = *currTile
						*currTile = 0
					} else if *nextTile == *currTile { // merge
						*nextTile = *currTile * 2
						*currTile = 0
                        break
					}
				}
			}
		}
	}
}

func moveRight(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := BOARD_SIZE - 1; j >= 0; j-- {
			if (*b)[i][j] != 0 {
				for pos := j; pos != BOARD_SIZE-1; pos++ {
					currTile := &(*b)[i][pos]
					nextTile := &(*b)[i][pos+1]

					if *nextTile == 0 { // slide
						*nextTile = *currTile
						*currTile = 0
					} else if *nextTile == *currTile { // merge
						*nextTile = *currTile * 2
						*currTile = 0
                        break
					}
				}
			}
		}
	}
}
