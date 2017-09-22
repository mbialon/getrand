package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/dustin/go-humanize"
)

func main() {
	size := "1KiB"
	if len(os.Args) > 1 {
		size = os.Args[1]
	}
	n, err := humanize.ParseBytes(size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse size, err: %v\n", err)
		os.Exit(1)
	}
	if _, err := io.CopyN(os.Stdout, rand.Reader, int64(n)); err != nil {
		fmt.Fprintf(os.Stderr, "cannot copy data: %v\n", err)
		os.Exit(1)
	}
}
