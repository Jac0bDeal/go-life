package output

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/buger/goterm"
)

// Terminal represents a terminal session and everything that can be done to it.
type Terminal struct {
}

// NewTerminal builds and returns a cbreak terminal session.
func NewTerminal() *Terminal {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("unable to activate cbreak mode: %v", err)
	}

	terminal := &Terminal{}
	terminal.clearScreen()

	return terminal
}

// Close changes the Terminal back to cooked mode.
func (t *Terminal) Close() error {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	return cookedTerm.Run()
}

// clearScreen wipes the screen by returning the cursor to the origin.
func (t *Terminal) clearScreen() {
	t.moveCursor(0, 0)
}

// moveCursor wraps goterm.MoveCursor to achieve zero indexing
// and maps slice indexing to goterm indexing.
func (t *Terminal) moveCursor(row, col int) {
	goterm.MoveCursor(col+1, row+1)
}

// Print prints the passed state to the Terminal.
func (t *Terminal) Print(state [][]int) error {
	t.clearScreen()
	width := len(state[0])
	height := len(state[1])

	var buffer bytes.Buffer
	buffer.WriteString("╔═" + strings.Repeat("═", width*2) + "═╗\n")
	for j := 0; j < height; j++ {
		buffer.WriteString("║ ")
		for i := 0; i < width; i++ {
			if state[i][j] == 1 {
				buffer.WriteString("██")
			} else {
				buffer.WriteString("  ")
			}
		}
		buffer.WriteString(" ║\n")
	}

	buffer.WriteString("╚═" + strings.Repeat("═", width*2) + "═╝\n")

	_, err := goterm.Print(buffer.String())
	if err != nil {
		return err
	}

	goterm.Flush()
	return nil
}
