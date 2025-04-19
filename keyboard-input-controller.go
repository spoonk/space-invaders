package main

import (
	// "errors"
	// "fmt"
	"time"
)

// todo: maybe some rising edge timer??
//
//	if we just started pressing a new key, repeat that key for 200ms?
//	cuz the keyboard needs time to warm up...

const TIMER_DURATION_MS = 100 * NANOSECOND
const NO_INPUT = -1

type KeyboardInputController struct {
	lastPressedKey   rune
	currentKeyPress  rune
	pressExpireTimer *time.Timer
}

func (k *KeyboardInputController) init(handler *KeyboardInputHandler) {
	for _, key := range []rune{'w', 'a', 's', 'd'} {
		handler.registerCallback(key, k.onKeypressReceive)
	}
}

func (k *KeyboardInputController) waitForTimer() {
	<-k.pressExpireTimer.C
	k.currentKeyPress = NO_INPUT
}

func (k *KeyboardInputController) onKeypressReceive(char rune) {
	isExistingWaiter := k.pressExpireTimer.Stop() // should be a thread waiting for this timer already
	k.pressExpireTimer.Reset(TIMER_DURATION_MS)
	if !isExistingWaiter {
		go k.waitForTimer()
	}

	k.currentKeyPress = char
	k.lastPressedKey = char
}

func (k *KeyboardInputController) getLastKeypress() rune {
	return k.lastPressedKey
	// return '0', errors.New("No last pressed key")
}

func (k *KeyboardInputController) getCurrentKeypress() rune {
	return k.currentKeyPress
	// return '0', errors.New("No key currently pressed")
}

func NewKeyBoardInputController() *KeyboardInputController {
	timer := time.NewTimer(0)
	timer.Stop()
	return &KeyboardInputController{pressExpireTimer: timer}
}
