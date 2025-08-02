package keyboard

import (
	"errors"
	"fmt"
	"golang.org/x/term"
	"os"
)

type KeyboardInputHandler struct {
	oldTerminalState *term.State
	inputFile        *os.File
	callbacks        map[rune][]func(rune)
}

func (k *KeyboardInputHandler) Init() {
	fd := int(k.inputFile.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

	k.oldTerminalState = oldState
}

func NewKeyboardInput() *KeyboardInputHandler {
	return &KeyboardInputHandler{inputFile: os.Stdin, callbacks: make(map[rune][]func(rune))}
}

func (k *KeyboardInputHandler) RegisterCallback(char rune, fun func(rune)) {
	k.callbacks[char] = append(k.callbacks[char], fun)
}

func (k *KeyboardInputHandler) Cleanup() {
	term.Restore(int(k.inputFile.Fd()), k.oldTerminalState)
}

func (k *KeyboardInputHandler) readInput() (byte, error) {
	// note: arrow keys will actually return 3 bytes, will need to handle those
	buffer := make([]byte, 1)

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

	return '0', fmt.Errorf("no registered mapping for byte %b", b)
}

func (k *KeyboardInputHandler) fireEventsForChar(char rune) {
	callbacks := k.callbacks[char]
	for _, callback := range callbacks {
		callback(char)
	}
}

func (k *KeyboardInputHandler) Loop() {
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
