package terminal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Terminal wraps terminal input and output behind a small reusable API.
type Terminal struct {
	in  *bufio.Reader
	out io.Writer
}

// NewTerminal builds a terminal helper from any reader and writer.
func NewTerminal(in io.Reader, out io.Writer) *Terminal {
	return &Terminal{
		in:  bufio.NewReader(in),
		out: out,
	}
}

// Println writes a line to the terminal.
func (t *Terminal) Println(args ...any) error {
	_, err := fmt.Fprintln(t.out, args...)
	return err
}

// Printf writes formatted output to the terminal.
func (t *Terminal) Printf(format string, args ...any) error {
	_, err := fmt.Fprintf(t.out, format, args...)
	return err
}

// ReadLine prints a prompt and returns the next line from input.
func (t *Terminal) ReadLine(prompt string) (string, error) {
	if prompt != "" {
		if _, err := fmt.Fprint(t.out, prompt); err != nil {
			return "", err
		}
	}

	line, err := t.in.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	if err != nil {
		if err == io.EOF && line != "" {
			return line, nil
		}
		return "", err
	}

	return line, nil
}
