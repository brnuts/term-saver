package main

import (
	"fmt"
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

const message = "Windows size is %dx%d"
const wait = time.Second / 100000

func printCurrentSize(x int, y int) {
	tm.MoveCursor(x, y)
	tm.Printf(message, tm.Width(), tm.Height())
}

func moveHorizontallyToRight(y int) {
	for x := 2; x <= tm.Width(); x++ {
		tm.MoveCursor(x, y)
		tm.Printf("*")
		time.Sleep(wait)
		tm.MoveCursor(x-1, y)
		tm.Printf(" ")
		tm.Flush()
	}
}

func moveVerticallyToDown(x int) {
	for y := 2; x <= tm.Height(); x++ {
		tm.MoveCursor(x, y)
		tm.Printf("*")
		time.Sleep(wait)
		tm.MoveCursor(x, y-1)
		tm.Printf(" ")
		tm.Flush()
	}
}

func randomWriteCharTerminal(char string) {
	x := rand.Intn(tm.Width())
	y := rand.Intn(tm.Height())
	tm.MoveCursor(x, y)
	//tm.Printf(char)
	colour256Code := fmt.Sprintf("%d", rand.Intn(16)*16+rand.Intn(16))
	colourFmt := "\033[38;5;" + colour256Code + "m %s"
	tm.Printf(colourFmt, char)
}

func main() {
	tm.Clear() // Clear current screen
	for y := 1; y <= 300000; y++ {
		//height := tm.Height()
		//width := tm.Width()
		randomWriteCharTerminal("+")
		//printCurrentSize(int(width-len(message)/2), int((height-1)/2))
		//moveHorizontallyToRight(y)
		tm.Flush()
		//time.Sleep(wait)
	}
	// Reset colours
	tm.MoveCursor(0, 0)
	tm.Println("\033[0m")
	tm.Flush()
	tm.Clear()

}
