package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
)

var SLEEP_DUR = 10 * NANOSECOND

// todo: loop in background
type KeyboardInputController struct {
	reader bufio.Reader
}

func (k *KeyboardInputController) refresh() {

	var b []byte = make([]byte, 1)

	_, err := os.Stdin.Read(b)

	if err == nil {
		keyPress <- rune(b[0])
	} else {
		print(err)
	}
}

func (k *KeyboardInputController) init() {
	fd := int(os.Stdin.Fd())
	_, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

}

func (k *KeyboardInputController) refreshEternally() {
	for {
		// k.init()
		k.refresh()
	}

}

func newKeyBoardInputController() *KeyboardInputController {
	return &KeyboardInputController{reader: *bufio.NewReader(os.Stdin)}
}
