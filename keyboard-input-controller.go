package main

import (
	"bufio"
	"os"
)

type KeyboardInputController struct {
	reader  bufio.Reader
	lastKey byte
}

func (k *KeyboardInputController) refresh() {
	byte, err := k.reader.ReadByte()
	if err != nil {
		k.lastKey = byte
	}
}

func newKeyBoardInputController() *KeyboardInputController {
	return &KeyboardInputController{reader: *bufio.NewReader(os.Stdin)}
}
