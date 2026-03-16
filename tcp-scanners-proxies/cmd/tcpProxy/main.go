package main

import (
	"fmt"
	"log"
	"os"
)

// Building a TCP Proxy
// First let's start with a echo server
// how the interfaces defined in go:
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

type FooReader struct{}
type FooWriter struct{}

func (FooReader *FooReader) Read(b []byte) (n int, err error) {
	// Read some data from anywhere.
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

func (FooWriter *FooWriter) Write(b []byte) (n int, err error) {
	// Write the data some where.
	fmt.Print("out >")
	return os.Stdout.Write(b)

}

func main() {
	// Instantiate reader and writer.
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output.
	input := make([]byte, 4096)

	// Use reader to read input.
	s, err := reader.Read(input)

	if err != nil {
		log.Fatalln("Unable to read data!")
	}

	fmt.Printf("Read %d bytes from stdio\n", s)

	// Use writer to write output.
	s, err = writer.Write(input)

	if err != nil {
		log.Fatalln("Unable to write data!")
	}

	fmt.Printf("Wrote %d bytes to stdout\n", s)
}

// func Copy(dst io.Writer, src io.Reader) (written int64, err error)

// func main() {
// 	var (
// 		reader FooReader
// 		writer FooWriter
// 	)

// 	if _, err := io.Copy(&writer, &reader); err != nil {
// 		log.Fatalln("Unable to read/write data")
// 	}
// }
