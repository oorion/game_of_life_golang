package display

import (
	"game_of_life/board"
)

func PrintBoard(b board.Board) string {
	var output string
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if b.Cells[i*b.Width+j].CurrentStatus {
				output += "1"
			} else {
				output += "0"
			}
		}
		output += "\n"
	}
	return output
}
