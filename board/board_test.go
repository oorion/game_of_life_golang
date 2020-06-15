package board

import (
	"game_of_life/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

var successTestCases = []struct {
	name     string
	offset   node.Position
	position node.Position
	height   int
	width    int
	expected int
}{
	{
		name:     "middle node upper left neighbor",
		offset:   node.Position{-1, -1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 0,
	},
	{
		name:     "middle node upper neighbor",
		offset:   node.Position{0, -1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 1,
	},
	{
		name:     "middle node upper right neighbor",
		offset:   node.Position{1, -1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 2,
	},
	{
		name:     "middle node left neighbor",
		offset:   node.Position{-1, 0},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 3,
	},
	{
		name:     "middle node right neighbor",
		offset:   node.Position{1, 0},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 5,
	},
	{
		name:     "middle node lower left neighbor",
		offset:   node.Position{-1, 1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 6,
	},
	{
		name:     "middle node lower neighbor",
		offset:   node.Position{0, 1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 7,
	},
	{
		name:     "middle node lower right neighbor",
		offset:   node.Position{1, 1},
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: 8,
	},
}

func TestComputeNeighborIndexSuccessCases(t *testing.T) {
	for _, tt := range successTestCases {
		actual, err := computeNeighborIndex(tt.height, tt.width, tt.offset, tt.position)
		assert.Nil(t, err, tt.name)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}

var errorTestCases = []struct {
	name     string
	offset   node.Position
	position node.Position
	height   int
	width    int
	expected int
}{
	{
		name:     "out of bounds to the left",
		offset:   node.Position{-1, -1},
		position: node.Position{0, 1},
		height:   3,
		width:    3,
	},
	{
		name:     "out of bounds to the right",
		offset:   node.Position{1, 0},
		position: node.Position{2, 1},
		height:   3,
		width:    3,
	},
	{
		name:     "out of bounds on top",
		offset:   node.Position{0, -1},
		position: node.Position{1, 0},
		height:   3,
		width:    3,
	},
	{
		name:     "out of bounds on top",
		offset:   node.Position{0, 1},
		position: node.Position{1, 2},
		height:   3,
		width:    3,
	},
}

func TestComputeNeighborIndexErrorCases(t *testing.T) {
	for _, tt := range errorTestCases {
		_, err := computeNeighborIndex(tt.height, tt.width, tt.offset, tt.position)
		assert.NotNil(t, err, tt.name)
	}
}

var successTestCasesIndices = []struct {
	name     string
	position node.Position
	height   int
	width    int
	expected []int
}{
	{
		name:     "middle node",
		position: node.Position{1, 1},
		height:   3,
		width:    3,
		expected: []int{
			0,
			1,
			2,
			3,
			5,
			6,
			7,
			8,
		},
	},
	{
		name:     "upper left corner node",
		position: node.Position{0, 0},
		height:   3,
		width:    3,
		expected: []int{
			1,
			3,
			4,
		},
	},
}

func TestComputeNeighborIndicesSuccessCases(t *testing.T) {
	for _, tt := range successTestCasesIndices {
		actual := computeNeighborIndices(tt.height, tt.width, tt.position)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}

var successTestCasesDeadBoard = []struct {
	name     string
	height   int
	width    int
	expected []node.Cell
}{
	{
		name:     "2 X 2",
		height:   2,
		width:    2,
		expected: expected2by2DeadBoard(),
	},
}

func expected2by2DeadBoard() []node.Cell {
	var board = make([]node.Cell, 4)
	board[0] = node.Cell{
		CurrentStatus: false,
		NextStatus:    false,
		Neighbors: []*node.Cell{
			&board[1],
			&board[2],
			&board[3],
		},
		Position: node.Position{
			X: 0,
			Y: 0,
		},
	}
	board[1] = node.Cell{
		CurrentStatus: false,
		NextStatus:    false,
		Neighbors: []*node.Cell{
			&board[0],
			&board[2],
			&board[3],
		},
		Position: node.Position{
			X: 1,
			Y: 0,
		},
	}
	board[2] = node.Cell{
		CurrentStatus: false,
		NextStatus:    false,
		Neighbors: []*node.Cell{
			&board[0],
			&board[1],
			&board[3],
		},
		Position: node.Position{
			X: 0,
			Y: 1,
		},
	}
	board[3] = node.Cell{
		CurrentStatus: false,
		NextStatus:    false,
		Neighbors: []*node.Cell{
			&board[0],
			&board[1],
			&board[2],
		},
		Position: node.Position{
			X: 1,
			Y: 1,
		},
	}

	return board
}

func TestInitializeDeadBoard(t *testing.T) {
	for _, tt := range successTestCasesDeadBoard {
		actual := InitializeDeadBoard(tt.height, tt.width)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}
