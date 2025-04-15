package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

type KeyboardInputController struct{}

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
		k.refresh()
	}
}

func newKeyBoardInputController() *KeyboardInputController {
	return &KeyboardInputController{}
}
