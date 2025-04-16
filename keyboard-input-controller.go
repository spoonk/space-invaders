package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"golang.org/x/term"
)

const NO_INPUT = -1

type KeyboardInputController struct {
	cnt int
}

func (k *KeyboardInputController) refresh() {
	var b []byte = make([]byte, 1)

	sz, err := os.Stdin.Read(b)

	if err == nil {
		if sz == 0 {
			kp = NO_INPUT
		} else {
			kp = rune(b[0])
		}
	} else {
		// assume EAGAIN lol
		// EAGAIN https://stackoverflow.com/questions/4058368/what-does-eagain-mean
		kp = NO_INPUT
	}
}

func (k *KeyboardInputController) init() {
	fd := int(os.Stdin.Fd())
	_, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

	err = syscall.SetNonblock(fd, true)
	if err != nil {
		print(err)
	}

}

func (k *KeyboardInputController) refreshEternally() {
	for {
		k.refresh()
		time.Sleep(NANOSECOND * 20)
	}
}

func newKeyBoardInputController() *KeyboardInputController {
	return &KeyboardInputController{cnt: 0}
}
