package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var successTestCases = []struct {
	name          string
	neighbors     []*Cell
	currentStatus bool
	nextStatus    bool
	expected      bool
}{
	{
		name: "alive with zero active neighbors should live",
		neighbors: []*Cell{
			&Cell{
				CurrentStatus: false,
			},
		},
		currentStatus: true,
		nextStatus:    true,
		expected:      false,
	},
	{
		name: "alive with 2 active neighbors should live",
		neighbors: []*Cell{
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
		},
		currentStatus: true,
		nextStatus:    false,
		expected:      true,
	},
	{
		name: "alive with 4 active neighbors should die",
		neighbors: []*Cell{
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
		},
		currentStatus: true,
		nextStatus:    true,
		expected:      false,
	},
	{
		name: "dead with 3 active neighbors should resurrect",
		neighbors: []*Cell{
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
			&Cell{
				CurrentStatus: true,
			},
		},
		currentStatus: false,
		nextStatus:    false,
		expected:      true,
	},
}

func TestUpdateNextStatus(t *testing.T) {
	for _, tt := range successTestCases {
		cell := Cell{
			CurrentStatus: tt.currentStatus,
			NextStatus:    tt.nextStatus,
			Neighbors:     tt.neighbors,
		}
		cell.UpdateNextStatus()
		actual := cell.NextStatus
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}
