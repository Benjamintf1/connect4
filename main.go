package main

import (
	"fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
	board := makeInitialBoard(7, 6)
	reader := bufio.NewReader(os.Stdin)
  turn := 1
	for !gameOver(board) {
    printBoard(board);
    fmt.Printf("it is %s's turn \n", turnChar(turn))
		fmt.Printf("Pick a column from 0 to %d, or choose to rotate the board right(r) or left(l)\n", len(board)-1)
		line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    switch {
    case func() bool { num, err := strconv.ParseInt(line, 10, 64); if err != nil { return false }; return int(num) < len(board)}():
      column, _ := strconv.ParseInt(line, 10, 64)
      turn = dropPiece(&board, turn, int(column))
    case strings.Contains(line, "r") || strings.Contains(line, "l"):
      board = rotateBoard(board, strings.Contains(line, "r"))
    default:
      fmt.Printf("you entered %s, input invalid\n", line)
    }
	}
}

func makeInitialBoard(columns, rows int) [][]int {
	board := make([][]int, columns)
	for i := range board {
		board[i] = make([]int, rows)
	}
	return board
}

func rotateBoard(board [][]int, isRight bool) [][]int {
  newBoard := make([][]int, len(board[0]))
  for i := range newBoard {
    newBoard[i] = make([]int, len(board))
  }
  for j := 0; j < len(board); j++ {
    column := j
    if isRight {
      column = len(board)-1-j
    }
    for i := 0; i < len(board[0]); i++ {
       newColumn := i
       if isRight {
         newColumn = len(board[0])-1-i
       }
       if board[column][i] != 0 {
        dropPiece(&newBoard,board[column][i], newColumn)
       }
    }
  }
  return newBoard
}

func gameOver(board [][]int) bool {
	return false
}

func printBoard(board [][]int) {
  for j := 0; j < len(board[0]); j++ {
    for i := 0; i < len(board); i++ {
      fmt.Print(turnChar(board[i][j]) + " ")
    }
    fmt.Print("\n")
  }
}

func turnChar(turn int) string {
  character := "0"
  if turn == 1 {
    character = "R"
  } else if turn == 2 {
    character = "B"
  }
  return character
}

func dropPiece(board *[][]int, turn, column int) int {
  for i := len((*board)[0])-1; i >= 0; i-- {
    if (*board)[column][i] == 0 {
      (*board)[column][i] = turn
      if turn == 1 {
        return 2
      }
      if turn == 2 {
        return 1
      }
    }
  }
  return turn
}
