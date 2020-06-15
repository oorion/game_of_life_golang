package main

import (
	"fmt"
	"game_of_life/board"
	"game_of_life/display"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	file, err := os.Open("glider1.txt")
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	var w int
	for _, element := range data[:count] {
		if element == '0' || element == '1' {
			w++
		} else {
			break
		}
	}
	h := count / (w + 1)
	b := board.Board{
		Height: h,
		Width:  w,
		Cells:  board.InitializeDeadBoard(h, w),
	}
	i := 0
	for _, element := range data[:count] {
		if element == '0' {
			b.Cells[i].CurrentStatus = false
			i++
		}
		if element == '1' {
			b.Cells[i].CurrentStatus = true
			i++
		}
	}

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println(display.PrintBoard(b))

		for i := range b.Cells {
			b.Cells[i].UpdateNextStatus()
		}
		for i := range b.Cells {
			b.Cells[i].CurrentStatus = b.Cells[i].NextStatus
		}
		time.Sleep(1 * time.Second)
	}
}
