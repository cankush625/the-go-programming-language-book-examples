package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		fmt.Println(r)
	}
	return nil
}

func main() {
	err := read()
	if err != nil {
		return
	}
}
