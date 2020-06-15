package node

type Cell struct {
	CurrentStatus bool
	NextStatus    bool
	Neighbors     []*Cell
	Position      Position
}

type Position struct {
	X int
	Y int
}

func (c *Cell) UpdateNextStatus() {
	var countOfLiveNeighbors int
	for _, neighbor := range c.Neighbors {
		if neighbor.CurrentStatus {
			countOfLiveNeighbors++
		}
	}
	if c.CurrentStatus && countOfLiveNeighbors < 2 {
		c.NextStatus = false
		return
	}
	if c.CurrentStatus && (countOfLiveNeighbors == 2 || countOfLiveNeighbors == 3) {
		c.NextStatus = true
		return
	}
	if c.CurrentStatus && countOfLiveNeighbors > 3 {
		c.NextStatus = false
		return
	}
	if !c.CurrentStatus && countOfLiveNeighbors == 3 {
		c.NextStatus = true
		return
	}
}
