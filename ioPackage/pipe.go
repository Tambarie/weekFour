package main

import (
	"bytes"
	"fmt"
	"io"
)

func main()  {
	r,w := io.Pipe()
	go func() {
		fmt.Fprint(w,"here are some files")
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Print(buf.String())
}