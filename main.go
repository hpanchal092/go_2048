package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// size of the board, always a square, ex: 4 means 4x4
const BOARD_SIZE int = 4

func main() {
	// randomize the seed because it isn't random by default
	seed := time.Now().Unix()
	rand.Seed(seed)

	fmt.Println("Welcome to GO 2048!")

	gameBoard := [BOARD_SIZE][BOARD_SIZE]int{}

	// game loop omg like python games course 😱
	for {
		addTile(&gameBoard)
		printBoard(&gameBoard)

        moves := checkValidMoves(&gameBoard)
        if (len(moves) == 0) {
            fmt.Println("you lost 🤡")
            time.Sleep(time.Second * 3)
            return
        }
		input := getInput(&moves)
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
	chance := rand.Intn(10)
	if chance == 0 {
		val = 4
	}

	// assign the value to the tile
	*tile = val

	return true
}

// move code below to separate file eventually, or not idk
func checkValidMoves(b *[BOARD_SIZE][BOARD_SIZE]int) []string {
    output := make([]string, 0, 4)

    uB := moveUp(*b)
    dB := moveDown(*b)
    lB := moveLeft(*b)
    rB := moveRight(*b)
    
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

func move(input string, b *[BOARD_SIZE][BOARD_SIZE]int) {
	switch input {
	case "u":
        *b = moveUp(*b)
	case "d":
        *b = moveDown(*b)
	case "l":
        *b = moveLeft(*b)
	case "r":
        *b = moveRight(*b)
	}
}

func moveUp(b [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (b)[i][j] != 0 {
				for pos := i; pos != 0; pos-- {
					currTile := &b[pos][j]
					nextTile := &b[pos-1][j]

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
	return b
}

func moveDown(b [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	// don't ask me how it works cuz idk
	for i := BOARD_SIZE - 1; i >= 0; i-- {
		for j := 0; j < BOARD_SIZE; j++ {
			if b[i][j] != 0 {
				for pos := i; pos != BOARD_SIZE-1; pos++ {
					currTile := &b[pos][j]
					nextTile := &b[pos+1][j]

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
	return b
}

func moveLeft(b [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b[i][j] != 0 {
				for pos := j; pos != 0; pos-- {
					currTile := &b[i][pos]
					nextTile := &b[i][pos-1]

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
	return b
}

func moveRight(b [BOARD_SIZE][BOARD_SIZE]int) [BOARD_SIZE][BOARD_SIZE]int {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := BOARD_SIZE - 1; j >= 0; j-- {
			if b[i][j] != 0 {
				for pos := j; pos != BOARD_SIZE-1; pos++ {
					currTile := &b[i][pos]
					nextTile := &b[i][pos+1]

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
	return b
}
