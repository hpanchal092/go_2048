package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// size of the board, always a square, ex: 4 means 4x4
const SIZE int = 4

func main() {
	// randomize the seed because it isn't random by default
	seed := time.Now().Unix()
	rand.Seed(seed)

	fmt.Println("Welcome to GO 2048!")

	gameBoard := [SIZE][SIZE]int{}

	// game loop omg like python games course 😱
	for {
		addTile(&gameBoard)
		printBoard(&gameBoard)

		moves := checkValidMoves(&gameBoard)
		if len(moves) == 0 {
			fmt.Println("you lost 🤡")
			time.Sleep(time.Second * 3)
			return
		}
		input := getInput(&moves)
		move(input, &gameBoard)
	}
}

func printBoard(b *[SIZE][SIZE]int) {
	// prints the board i dont know what else to say
	fmt.Printf("\n")
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getInput(moves *[]string) string {
	// gets input from the user, either u, d, l, or r
	var userInput string

	for !isValid(&userInput, moves) {
		fmt.Printf("Enter an input (u/d/l/r): ")
		fmt.Scanf("%s", &userInput)
	}

	return userInput
}

func isValid(s *string, moves *[]string) bool {
	// takes in a string s and checks if it is a valid input, returns a bool
	// valid inputs start with the letter u, d, l, or r (case insensitive)
	// ignores inputs if the move is invalid aka not in moves slice

	if len(*s) == 0 {
		return false
	}
	*s = strings.ToLower(*s)
	*s = (*s)[:1]
	for i := 0; i < len(*moves); i++ {
		if *s == (*moves)[i] {
			return true
		}
	}
	fmt.Println("Invalid input entered")
	return false
}

func addTile(b *[SIZE][SIZE]int) bool {
	// takes in the board and returns true if it adds a tile successfully
	// returns false if unsuccessful aka the board is full

	// create a slice of all of the empty tiles (tiles with a value of 0)
	emptyTiles := make([]*int, 0, 16)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
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
	chance := rand.Intn(10)
	if chance == 0 {
		val = 4
	}

	// assign the value to the tile
	*tile = val

	return true
}

func checkValidMoves(b *[SIZE][SIZE]int) []string {
	output := make([]string, 0, 4)

	// generate the 4 possible board positions from the current position
	uB := *b
	move("u", &uB)

	dB := *b
	move("d", &dB)

	lB := *b
	move("l", &lB)

	rB := *b
	move("r", &rB)

	// add boards that changed to the list of valid moves
	if *b != uB {
		output = append(output, "u")
	}
	if *b != dB {
		output = append(output, "d")
	}
	if *b != lB {
		output = append(output, "l")
	}
	if *b != rB {
		output = append(output, "r")
	}

	return output
}

func move(input string, b *[SIZE][SIZE]int) {
	// creates a slice representing a row/column on the board, then slides and
	// merges pieces on the row/column

	slice := []*int{}
	for i := 0; i < SIZE; i++ {
		switch input {
		case "u":
			slice = createVerticalSlice(b, i, true)
		case "d":
			slice = createVerticalSlice(b, i, false)
		case "l":
			slice = createHorizontalSlice(b, i, true)
		case "r":
			slice = createHorizontalSlice(b, i, false)
		}

		slideAndMerge(slice)
	}
}

func slideAndMerge(s []*int) {
	for i := 1; i < SIZE; i++ { // loop through the tiles
		if *s[i] != 0 { // if it has a proper value
			for pos := i; pos > 0; pos-- { // start sliding it towards the left
				currTile := s[pos]
				nextTile := s[pos-1]

				// merge if next tile is same as curr tile
				if *nextTile == *currTile {
					// makes it negative so other tiles can't merge with it in
					// the same move
					*nextTile = *currTile * -2
					*currTile = 0
					break
				}
				// stop sliding if next tile is not 0
				if *nextTile != 0 {
					break
				}
				// slide
				*nextTile = *currTile
				*currTile = 0
			}
		}
	}
	// remove the negative sign
	for i := 0; i < SIZE; i++ {
		*s[i] = int(math.Abs(float64(*s[i])))
	}
}

func createVerticalSlice(b *[SIZE][SIZE]int, n int, forwards bool) []*int {
	slice := make([]*int, 0, 4)

	if forwards {
		// create slice left, nth row
		for i := 0; i < SIZE; i++ {
			slice = append(slice, &b[i][n])
		}
	} else {
		// create slice right, nth row
		for i := SIZE - 1; i >= 0; i-- {
			slice = append(slice, &b[i][n])
		}
	}
	return slice
}

func createHorizontalSlice(b *[SIZE][SIZE]int, n int, forwards bool) []*int {
	slice := make([]*int, 0, 4)

	if forwards {
		// create slice upwards, nth column
		for i := 0; i < SIZE; i++ {
			slice = append(slice, &b[n][i])
		}
	} else {
		// create slice downwards, nth column
		for i := SIZE - 1; i >= 0; i-- {
			slice = append(slice, &b[n][i])
		}
	}
	return slice
}
