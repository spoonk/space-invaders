package main

import (
	"errors"
	"fmt"
	"time"
)

const NO_INPUT = -1

type KeyboardInputController struct {
	lastPressedKey  rune
	currentKeyPress rune
	ticket          *time.Timer
}

func (k *KeyboardInputController) init(handler *KeyboardInputHandler) {
	for _, key := range []rune{'w', 'a', 's', 'd'} {
		handler.registerCallback(key, k.callbackFunction)
	}
}

func (k *KeyboardInputController) callbackFunction(char rune) {
	// set/reset timer here
	fmt.Printf("you outputted: %c\n", char)
}

func (k *KeyboardInputController) getLastKeypress() (rune, error) {
	return '0', errors.New("No last pressed key")
}

func (k *KeyboardInputController) getCurrentKeypress() (rune, error) {
	return '0', errors.New("No key currently pressed")
}

func NewKeyBoardInputController() *KeyboardInputController {
	return &KeyboardInputController{}
}
