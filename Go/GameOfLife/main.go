package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board struct {
	size       int
	boardArray [][]int
}

func readInput() int {
	var size int
	fmt.Scanln(&size)
	return size
}
func (b *Board) Initialize() {
	for i := 0; i < b.size; i++ {
		b.boardArray[i] = make([]int, b.size)
		for j := 0; j < b.size; j++ {
			b.boardArray[i][j] = rand.Intn(2)
		}
	}
}
func createRandomBoard(size int) Board {
	board := Board{size: size, boardArray: make([][]int, size)}
	board.Initialize()
	return board
}
func (b *Board) GetFormattedBoard() string {
	var builder strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.boardArray[i][j] == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteString("O")
			}
			if j == b.size-1 {
				builder.WriteString("\n")
			}
		}
	}
	return builder.String()
}
func (b *Board) getLiveNeighbors(i int, j int) int {
	var neighbors, newI, newJ int
	di := [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	dj := [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	for k := 0; k < 8; k++ {
		newI = i + di[k]
		newJ = j + dj[k]
		if newI < 0 {
			newI = b.size - 1
		}
		if newI == b.size {
			newI = 0
		}
		if newJ < 0 {
			newJ = b.size - 1
		}
		if newJ == b.size {
			newJ = 0
		}
		// fmt.Println("my_value: ", b.GetCell(i, j), "i:", i, "j:", j, "newI:", newI, "newJ:", newJ, "value:", b.GetCell(newI, newJ))
		neighbors += b.GetCell(newI, newJ)
	}
	return neighbors
}

func (b *Board) SetCell(i, j, value int) {
	b.boardArray[i][j] = value
}
func (b *Board) GetCell(i, j int) int {
	return b.boardArray[i][j]
}
func (b *Board) NextGen() {
	var newBoard [][]int = make([][]int, b.size)
	for i := 0; i < b.size; i++ {
		newBoard[i] = make([]int, b.size)
	}
	var liveNeighbors, cellValue int
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			liveNeighbors = b.getLiveNeighbors(i, j)
			cellValue = b.GetCell(i, j)
			if cellValue == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				// fmt.Printf("setting cell i=%d, j=%d to zero because it has %d neighbors\n", i, j, liveNeighbors)
				newBoard[i][j] = 0
			} else if cellValue == 1 && !(liveNeighbors < 2 || liveNeighbors > 3) {
				newBoard[i][j] = 1
			} else if cellValue == 0 && liveNeighbors == 3 {
				// fmt.Printf("setting cell i=%d, j=%d to one because it has %d neighbors\n", i, j, liveNeighbors)
				newBoard[i][j] = 1
			} else {
				newBoard[i][j] = 0
			}
		}
	}
	b.boardArray = newBoard
}
func (b *Board) getNumberOfAliveCells() int {
	var aliveCells int
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			aliveCells += b.GetCell(i, j)
		}
	}
	return aliveCells
}
func main() {
	var size int = readInput()
	var board Board = createRandomBoard(size)
	for i := 0; i < 12; i++ {
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", board.getNumberOfAliveCells())
		fmt.Print(board.GetFormattedBoard())
		board.NextGen()
	}
}
