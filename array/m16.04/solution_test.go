package m16_04

import (
	"fmt"
	"testing"
)

func TestTictactoe(t *testing.T) {
	board := []string{" O    X", " O     ", "     O ", "XXXXXXX", "  O    ", " O   O ", "O  X OO"}
	fmt.Println(tictactoe(board))
}
