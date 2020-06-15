package board

import (
	"errors"
	"game_of_life/node"
)

type Board struct {
	Height int
	Width  int
	Cells  []node.Cell
}

// InitializeBoard generate a slice of blank cells that know about their neighbors and position
func InitializeDeadBoard(height int, width int) []node.Cell {
	var board = make([]node.Cell, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			position := node.Position{
				X: x,
				Y: y,
			}
			neighborIndices := computeNeighborIndices(height, width, position)
			neighbors := []*node.Cell{}
			for _, i := range neighborIndices {
				neighbors = append(neighbors, &board[i])
			}
			board[index] = node.Cell{
				CurrentStatus: false,
				NextStatus:    false,
				Neighbors:     neighbors,
				Position:      position,
			}
		}
	}

	return board
}

func computeNeighborIndices(height int, width int, position node.Position) []int {
	offsets := []node.Position{
		node.Position{-1, -1},
		node.Position{0, -1},
		node.Position{1, -1},
		node.Position{-1, 0},
		node.Position{1, 0},
		node.Position{-1, 1},
		node.Position{0, 1},
		node.Position{1, 1},
	}

	indices := []int{}
	for _, offset := range offsets {
		if index, err := computeNeighborIndex(height, width, offset, position); err == nil {
			indices = append(indices, index)
		}
	}
	return indices
}

func computeNeighborIndex(height, width int, offset, position node.Position) (int, error) {
	x := offset.X + position.X
	if x < 0 || x >= width {
		return 0, errors.New("out of bounds")
	}
	y := offset.Y + position.Y
	if y < 0 || y >= height {
		return 0, errors.New("out of bounds")
	}
	return y*width + x, nil
}
