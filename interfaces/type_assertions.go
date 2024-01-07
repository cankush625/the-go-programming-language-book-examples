package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success: ok, f == os.Stdout
	fmt.Println(f, ok)
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Println(b, ok)
}
