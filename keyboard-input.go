package main

// todo: eventually I'd like to replace this library with actual implementations to figure it out
import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

// detect keypress, fire event when one is pressed

type KeyboardInputHandler struct {
	oldTerminalState *term.State
	inputFile        *os.File
	callbacks        map[rune][]func(rune)
}

func (k *KeyboardInputHandler) init() {
	fd := int(k.inputFile.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

	k.oldTerminalState = oldState
	k.registerCallback('q', func(_ rune) { k.cleanup(); Run = false })
}

func NewKeyboardInput() *KeyboardInputHandler {
	return &KeyboardInputHandler{inputFile: os.Stdin, callbacks: make(map[rune][]func(rune))}
}

func (k *KeyboardInputHandler) registerCallback(char rune, fun func(rune)) {
	k.callbacks[char] = append(k.callbacks[char], fun)
}

func (k *KeyboardInputHandler) cleanup() {
	term.Restore(int(k.inputFile.Fd()), k.oldTerminalState)
}

func (k *KeyboardInputHandler) readInput() (byte, error) {
	// note: arrow keys will actually return 3 bytes, will need to handle those
	var buffer []byte = make([]byte, 1)

	_, err := os.Stdin.Read(buffer)
	if err != nil {
		return 0, errors.New("bad input")
	}

	return buffer[0], nil
}

func byteToCharacter(b byte) (rune, error) {
	switch b {
	case 119:
		return 'w', nil
	case 97:
		return 'a', nil
	case 115:
		return 's', nil
	case 100:
		return 'd', nil
	case 32:
		return ' ', nil
	case 113:
		return 'q', nil
	case 114:
		return 'r', nil
	case 49:
		return '1', nil
	}

	return '0', fmt.Errorf("No registered mapping for byte %b", b)
}

func (k *KeyboardInputHandler) fireEventsForChar(char rune) {
	callbacks := k.callbacks[char]
	for _, callback := range callbacks {
		callback(char)
	}
}

func (k *KeyboardInputHandler) loop() {
	for {
		b, err := k.readInput()
		if err != nil {
			continue
		}
		char, err := byteToCharacter(b)
		if err != nil {
			continue
		}

		k.fireEventsForChar(char)
	}
}
