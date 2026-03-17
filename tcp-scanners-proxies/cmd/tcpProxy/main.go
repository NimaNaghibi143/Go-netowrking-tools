package main

import (
	"io"
	"log"
	"net"
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

// type FooReader struct{}
// type FooWriter struct{}

// func (FooReader *FooReader) Read(b []byte) (n int, err error) {
// 	// Read some data from anywhere.
// 	fmt.Print("in >")
// 	return os.Stdin.Read(b)
// }

// func (FooWriter *FooWriter) Write(b []byte) (n int, err error) {
// 	// Write the data some where.
// 	fmt.Print("out >")
// 	return os.Stdout.Write(b)

// }

// func main() {
// 	// Instantiate reader and writer.
// 	var (
// 		reader FooReader
// 		writer FooWriter
// 	)

// 	// Create buffer to hold input/output.
// 	input := make([]byte, 4096)

// 	// Use reader to read input.
// 	s, err := reader.Read(input)

// 	if err != nil {
// 		log.Fatalln("Unable to read data!")
// 	}

// 	fmt.Printf("Read %d bytes from stdio\n", s)

// 	// Use writer to write output.
// 	s, err = writer.Write(input)

// 	if err != nil {
// 		log.Fatalln("Unable to write data!")
// 	}

// 	fmt.Printf("Wrote %d bytes to stdout\n", s)
// }

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

// Creating a echo server:

// echo is a handler function that simply echoes received data.

func ech(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)

	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])

		if err == io.EOF {
			log.Println("Some shit went wrong in client side! probably the connection shit!")
			break
		}

		if err != nil {
			log.Println("Some unexpected shit just went wrong!")
			break
		}

		log.Printf("Recieved %d bytes: %s\n", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data!")
		}
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")

	if err != nil {
		log.Fatalln("Unable to bind to port!")
	}

	log.Println("Listening on 0.0.0.0:20080")

	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		// Handle the connection. Using goroutine for concurrency.
		go ech(conn)

	}
}
