package repl

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type CommandHistory struct {
	commands []string
	index    int
}

func NewCommandHistory() *CommandHistory {
	return &CommandHistory{
		commands: make([]string, 0),
		index:    -1,
	}
}

func (h *CommandHistory) Add(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd != "" && (len(h.commands) == 0 || h.commands[len(h.commands)-1] != cmd) {
		h.commands = append(h.commands, cmd)

	}
	h.index = len(h.commands)
}

func (h *CommandHistory) Up() string {
	if len(h.commands) == 0 {
		return ""
	}
	if h.index == len(h.commands) {
		h.index = len(h.commands) - 1
	} else if h.index > 0 {
		h.index--
	}
	return h.commands[h.index]
}

func (h *CommandHistory) Down() string {
	if len(h.commands) == 0 || h.index >= len(h.commands)-1 {
		h.index = len(h.commands)
		return ""
	}
	h.index++
	return h.commands[h.index]
}

func readLineWithHistory(prompt string, history *CommandHistory) (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	fmt.Print(prompt)

	var input []rune
	buf := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return "", err
		}
		if n == 0 {
			continue
		}

		b := buf[0]

		if b == 3 {
			fmt.Print("^C\r\n")
			return "", fmt.Errorf("interrupted")
		} else if b == 4 {
			if len(input) == 0 {
				return "", fmt.Errorf("EOF")
			}
			input = input[:len(input)-1]
			fmt.Print("\b \b")
		} else if b == 13 || b == 10 {
			// 13 = Carriage Return (CR), 10 = Line Feed (LF) - Enter key

			fmt.Print("\r\n")
			return string(input), nil
		} else if b == 27 {
			// 27 = ESC - Start of escape sequence (arrow keys send ESC [ A/B/C/D)
			seq := make([]byte, 2)
			os.Stdin.Read(seq)

			if seq[0] == '[' {
				// '['  ANSI escape code for arrow keys
				switch seq[1] {
				case 'A': // 'A' = Up arrow (ESC [ A)
					if cmd := history.Up(); cmd != "" {
						input = []rune(cmd)
						fmt.Print("\r\033[K" + prompt + cmd)
					}
				case 'B': // 'B' = Down arrow (ESC [ B)
					if cmd := history.Down(); cmd != "" {
						input = []rune(cmd)
						fmt.Print("\r\033[K" + prompt + cmd)
					}
				}
			}
		} else if b >= 32 && b < 127 {
			// printable ASCII characters
			input = append(input, rune(b))
			fmt.Printf("%c", b)
		} else if b == 127 {
			//  Backspace ASCII character
			if len(input) > 0 {
				input = input[:len(input)-1]
				fmt.Print("\b \b")
			}

		}

	}
}
