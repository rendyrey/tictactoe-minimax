package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getChar(i int) string {
	switch i {
	case -1:
		return "X" // for player (using -1 value)
	case 0:
		return " " // for empty node
	case 1:
		return "O" // for cpu (using 1 value)
	}
	return " "
}

func draw(b [9]int) {
	fmt.Printf(" %v | %v | %v\n", getChar(b[0]), getChar(b[1]), getChar(b[2]))
	fmt.Printf("---+---+---\n")
	fmt.Printf(" %v | %v | %v\n", getChar(b[3]), getChar(b[4]), getChar(b[5]))
	fmt.Printf("---+---+---\n")
	fmt.Printf(" %v | %v | %v\n", getChar(b[6]), getChar(b[7]), getChar(b[8]))
}

func win(board [9]int) int {
	wins := [8][3]uint{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}} // pattern to be win
	// checking all patterns, if the pattern has similar symbol, either O or X, return the value, -1 for X, 1 for O
	for i := 0; i < 8; i++ {
		if board[wins[i][0]] != 0 && board[wins[i][0]] == board[wins[i][1]] && board[wins[i][0]] == board[wins[i][2]] {
			return board[wins[i][0]]
		}
	}
	// return zero if no one has the correct pattern
	return 0
}

func minimax(board [9]int, player int) int {
	winner := win(board)
	if winner != 0 {
		return winner * player
	}
	move := -1
	score := -2
	for i := 0; i < 9; i++ { // loop through the available slot
		if board[i] == 0 {
			board[i] = player                       // if the slot is empty, insert the slot for player, -1 or 1
			thisScore := -minimax(board, player*-1) // recursively to the leaf node and get the biggest score, explore all the condition possible of the game
			if thisScore > score {
				score = thisScore
				move = i
			}
			board[i] = 0
		}
	}
	if move == -1 {
		return 0
	}
	// fmt.Println("Score:", score)
	return score
}

func cpuMove(board [9]int) [9]int {
	move := -1
	score := -2
	for i := 0; i < 9; i++ {
		if board[i] == 0 {
			board[i] = 1                     // explore all the empty slot, try every empty slot to be filled by 'O' (cpu)
			tempScore := -minimax(board, -1) // get the max score from all game condition possible
			board[i] = 0
			if tempScore > score {
				score = tempScore
				move = i // get the index on the max score possible
			}
			// fmt.Printf("tempScore di kotak %d:%d\n", i, tempScore)
		}
	}
	// fmt.Printf("Score akhir => %d\n", score)
	board[move] = 1 // filled the empty index
	return board
}

func playerMove(board [9]int) [9]int { // for user/player input
	move := -1
	for move >= 9 || move < 0 {
		fmt.Printf("\nInput move (index 0 to 8): ")
		fmt.Scanln(&move)
	}
	if board[move] != 0 {
		fmt.Println("Please choose the empty space")
		playerMove(board)
	} else {
		board[move] = -1
	}
	return board
}

func main() {
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("Computer: O, You: X\n")
	rand.Seed(time.Now().UnixNano()) // seeder for random
	player := rand.Intn(2) + 1       // random, either 1 or 2
	if player == 1 {
		fmt.Printf("You play on first turn.\n")
	} else {
		fmt.Printf("You play on second turn\n")
	}
	turn := 0
	firstCPU := rand.Intn(9)          // random index for cpu if on the first turn, because if not, cpu always choose 0 index.
	for turn < 9 && win(board) == 0 { // while winner has not been determined and hasn't reach 8 turns yet.
		if (turn+player)%2 == 0 {
			if turn == 0 { // if cpu first, insert the slot on the random index
				board[firstCPU] = 1
			} else {
				board = cpuMove(board)
			}
		} else {
			draw(board)
			board = playerMove(board)
		}
		turn++
	}

	switch win(board) {
	case 0:
		draw(board)
		fmt.Printf("Draw Game\n")
		break
	case 1:
		draw(board)
		fmt.Printf("You lose\n")
		break
	case -1:
		draw(board)
		fmt.Print("You WIN!\n")
		break
	}
}
