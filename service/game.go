package service

import (
	"errors"
	"math/rand"
	"log"
)

type Cell struct {
	Value int
	Status string
}

type Game struct {
	Cells [][]Cell
	IsOpen [][]bool
	Row, Col int
}

var game Game

// TODO: Handle sessions
func GetGame() Game {
	log.Println("GetGame")
	return game
}

func Start(row, col, difficulty int) {
	game = Game {
		Row: row,
		Col: col,
		Cells: makeFilledArr(row, col, Cell{0, "success"}),
		IsOpen: makeFilledArr(row, col, false),
	}

	log.Println("Intialized")

	initializeMap(difficulty)
}

func Reveal(i, j int) (*Cell, error) {
	if !inRange(i, j) {
		return nil, errors.New("Invalid move")
	}

	if game.IsOpen[i][j] {
		return nil, errors.New("Cell has already been opened")
	}

	game.IsOpen[i][j] = true
	return &game.Cells[i][j], nil
}

func inRange(i, j int) bool {
	return 0 <= i && i <= game.Row && 0 <= j && j <= game.Col;
}

func makeFilledArr[T any](n int, m int, value T) [][]T {
	arr := make([][]T, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]T, m)
		for j := 0; j < m; j++ {
			arr[i][j] = value
		}
	}
	return arr
}

func initializeMap(difficulty int) {
	bombNum := (game.Row * game.Col) * difficulty / 100

	for cnt := 0; cnt < bombNum; cnt++ {
		for {
			i := rand.Intn(game.Row)
			j := rand.Intn(game.Col)
			if game.Cells[i][j].Value >= 0 {
				game.Cells[i][j].Value = -1
				game.Cells[i][j].Status = "danger"

				for ii := max(i - 1, 0); ii < min(i + 2, game.Row); ii++ {
					for jj := max(j - 1, 0); jj < min(j + 2, game.Col); jj++ {
						if (ii == i && jj == j) || (game.Cells[ii][jj].Value == -1) {
							continue
						}

						game.Cells[ii][jj].Value += 1
					}
				}
				break
			}
		}
	}
}
