package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "runtime"
  "strconv"
)

type board struct {
  grid [][]string
  dim int
}

func newBoard(dim int) board {
  // Takes board dimension and creates a new board object with an empty grid
  g := make([][]string, dim)
  for i := 0; i < dim; i++ {
    g[i] = make([]string, dim)
  }
  b := board{grid: g, dim: dim}
  return b
}

func (b *board) resetGrid() {
  for i := 0; i < b.dim; i++ {
    for j := 0; j < b.dim; j++ {
      b.grid[i][j] = ""
    }
  }
}

func (b board) display() {
    for i := 0; i < b.dim; i++ {
      fmt.Println(b.grid[i])
    }
}

func readInput() string {
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  if runtime.GOOS == "windows" {
    input = strings.TrimRight(input, "\r\n")
  } else {
    input = strings.TrimRight(input, "\n")
  }
  return input
}

func main() {
  fmt.Println("Welcome to Go!\n")
  fmt.Println("Enter board dimensions: (9, 13, 19)")
  retry := true
  var input string = ""
  inp := &input
  for retry {
    input := readInput()
    *inp = input
    if input == "9" || input == "13" || input == "19" {
      retry = false
    } else {
      fmt.Println("Please enter a board dimension from 9, 13 or 19:")
    }
  }

  boardDim := 0
  if input == "9" {
    boardDim = 9
  } else if input == "13" {
    boardDim = 13
  } else if input == "19" {
    boardDim = 19
  } else {
    fmt.Println("Input error - didn't get input from accepted values")
  }

  fmt.Println("\nBlack stones will be represented on the board as B, white as W")
  fmt.Println("When asked to pick a point, put in the format 'width height'")
  fmt.Println("If you wish to pass your turn, input 'PASS'. ")
  fmt.Println("Two passes from each player in succession ends the game.")

  board := newBoard(boardDim)
  end := false
  blacksTurn := true
  previousMovePassed := false
  for !end {
    if blacksTurn {
      fmt.Println("Black, please pick a point to place your stone: ")
    } else {
      fmt.Println("White, please pick a point to place your stone: ")
    }
    input = readInput()
    if input == "PASS" {
      if previousMovePassed {
        end = true
        break
      } else {
        previousMovePassed = true
      }
    } else {
      previousMovePassed = false

      // Returns the x and y points as a 1x2 array
      p := strings.Fields(input)
      p1, err := strconv.Atoi(p[0])
      p2, err := strconv.Atoi(p[1])
      if err != nil {
        fmt.Println(err)
      }
      if blacksTurn {
        board.grid[p1][p2] = "B"
      } else {
        board.grid[p1][p2] = "W"
      }
    }
    board.display()
    blacksTurn = !blacksTurn
  }

}
