package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(data []byte) (int, error) {
	n, err := rd.r.Read(data)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(data); i++ {
		char := data[i]
		if (char >= 'a' && char <= 'm') || (char >= 'A' && char <= 'M') {
			data[i] += 13
		} else if (char >= 'n' && char <= 'z') || (char >= 'N' && char <= 'Z') {
			data[i] -= 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
