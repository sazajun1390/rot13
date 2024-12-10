package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	// Read data from the underlying reader into p
	n, err = r.r.Read(p)
	if err != nil {
		return n, err
	}

	// Apply ROT13 transformation
	for i := 0; i < n; i++ {
		p[i] = r.rot13Decode(p[i])
	}
	return n, nil
}

// rot13 function applies the ROT13 transformation to a single byte
func (r *rot13Reader) rot13Decode(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z': // For uppercase letters
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z': // For lowercase letters
		return 'a' + (b-'a'+13)%26
	default:
		return b // Non-alphabetic characters remain unchanged
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
