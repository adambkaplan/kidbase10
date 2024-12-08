package encoding

import (
	"fmt"
	"io"
)

var DefaultEncoding = map[rune]string{
	' ': "0",
	'A': "1",
	'E': "2",
	'H': "3",
	'I': "4",
	'L': "5",
	'N': "6",
	'R': "7",
	'S': "8",
	'T': "9",
}

var DefaultLineLength = 8

type Encoder struct {
	writer     io.Writer
	encoding   map[rune]string
	lineLength int
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer:     w,
		encoding:   DefaultEncoding,
		lineLength: DefaultLineLength,
	}
}

func (e *Encoder) Encode(data string) error {
	dataLen := len(data)
	for i, r := range data {
		encoded, ok := e.encoding[r]
		if !ok {
			return fmt.Errorf("invalid input character %q", r)
		}
		if _, err := e.writer.Write([]byte(encoded)); err != nil {
			return fmt.Errorf("failure writing data: %v", err)
		}
		// If character position is not at the end of the data, add new line after 8 characters
		charPos := i + 1
		if charPos < dataLen && charPos%e.lineLength == 0 {
			if _, err := e.writer.Write([]byte("\n")); err != nil {
				return fmt.Errorf("failure writing data: %v", err)
			}
		}
	}
	// Padding - default pad to 8 characters
	remainder := dataLen % e.lineLength
	if remainder == 0 {
		return nil
	}
	zeroPad := e.lineLength - remainder
	zeroVal := e.encoding[' ']
	for range zeroPad {
		if _, err := e.writer.Write([]byte(zeroVal)); err != nil {
			return fmt.Errorf("failure writing data: %v", err)
		}
	}
	return nil
}
