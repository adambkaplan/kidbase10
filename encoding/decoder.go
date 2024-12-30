package encoding

import (
	"fmt"
	"io"
	"strings"
)

var DefaultDecoding = map[rune]string{
	'0': " ",
	'1': "A",
	'2': "E",
	'3': "H",
	'4': "I",
	'5': "L",
	'6': "N",
	'7': "R",
	'8': "S",
	'9': "T",
}

type Decoder struct {
	reader     io.Reader
	decoding   map[rune]string
	lineLength int
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader:     r,
		decoding:   DefaultDecoding,
		lineLength: DefaultLineLength,
	}
}

func (d *Decoder) Decode() (string, error) {
	data, err := io.ReadAll(d.reader)
	if err != nil {
		return "", fmt.Errorf("failed to read data: %v", err)
	}
	dataString := string(data)
	outBuilder := &strings.Builder{}
	for _, r := range dataString {
		if r == '\n' {
			continue
		}
		decoded, ok := d.decoding[r]
		if !ok {
			return "", fmt.Errorf("invalid input character %q", r)
		}
		outBuilder.WriteString(decoded)
	}
	zeroVal := d.decoding['0']
	return strings.TrimRight(outBuilder.String(), zeroVal), nil
}

func (d *Decoder) DecodeTo(w io.Writer) (int, error) {
	out, err := d.Decode()
	if err != nil {
		return 0, err
	}
	return io.WriteString(w, out)
}
