package main

import (
	"bufio"
	"io"
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

// ********this is the first version of the echo server.**********
// func ech(conn net.Conn) {
// 	defer conn.Close()

// 	// Create a buffer to store received data.
// 	b := make([]byte, 512)

// 	for {
// 		// Receive data via conn.Read into a buffer.
// 		size, err := conn.Read(b[0:])

// 		if err == io.EOF {
// 			log.Println("Some shit went wrong in client side! probably the connection shit!")
// 			break
// 		}

// 		if err != nil {
// 			log.Println("Some unexpected shit just went wrong!")
// 			break
// 		}

// 		log.Printf("Recieved %d bytes: %s\n", size, string(b))

// 		// Send data via conn.Write.
// 		log.Println("Writing data")

// 		if _, err := conn.Write(b[0:size]); err != nil {
// 			log.Fatalln("Unable to write data!")
// 		}
// 	}
// }

// Enhanced version of the echo function
// func echo(conn net.Conn) {
// 	defer conn.Close()

// 	reader := bufio.NewReader(conn)

// 	log.Println("Reading data:")
// 	s, err := reader.ReadString('\n')

// 	if err != nil {
// 		log.Println("Unable to read data!")
// 	}

// 	log.Printf("Read %d bytes: %s", len(s), s)

// 	log.Println("Writing data:")

// 	writer := bufio.NewWriter(conn)

// 	if _, err := writer.WriteString(s); err != nil {
// 		log.Fatalln("Unable to write data")
// 	}
// 	writer.Flush()
// }

// func main() {
// 	// Bind to TCP port 20080 on all interfaces.
// 	listener, err := net.Listen("tcp", ":20080")

// 	if err != nil {
// 		log.Fatalln("Unable to bind to port!")
// 	}

// 	log.Println("Listening on 0.0.0.0:20080")

// 	for {
// 		// Wait for connection. Create net.Conn on connection established.
// 		conn, err := listener.Accept()
// 		log.Println("Received connection")

// 		if err != nil {
// 			log.Fatalln("Unable to accept connection")
// 		}

// 		// Handle the connection. Using goroutine for concurrency.
// 		go echo(conn)

// 	}
// }

// Proxying a TCP Client

// what we are trying to achieve now is:
// a simple port forwarder to proxy a connection through an intermediary service or host.

// func handler(src net.Conn) {
// 	defer src.Close()
// 	dst, err := net.Dial("tcp", "joescatcam.website:80")

// 	if err != nil {
// 		log.Println("Unable to connect to our unreachable host!")
// 	}

// 	defer dst.Close()

// 	// Run in goroutine to prevent io.Copy from blocking
// 	go func() {
// 		// Copy our source's output to the destination
// 		if _, err := io.Copy(dst, src); err != nil {
// 			log.Println(err)
// 		}
// 	}()

// 	// Copy our destination's output back to our source
// 	if _, err := io.Copy(src, dst); err != nil {
// 		log.Println(err)
// 	}
// }

// func main() {
// 	// Listen on local port 80
// 	listener, err := net.Listen("tcp", ":8080")

// 	if err != nil {
// 		log.Fatalln("Unable to bind to port")
// 	}

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Println("Unable to accept connection!")
// 			continue
// 		}

// 		go handler(conn)
// 	}
// }

// Replicating Netcat for Command Execution
// $ nc –lp 13337 –e /bin/bash
// This command creates a listening server on port 13337. Any remote
// client that connects, perhaps via Telnet, would be able to execute arbitrary
// bash commands—hence the reason this is referred to as a gaping security hole.
// Netcat allows you to optionally include this feature during program compilation.

// Flusher wraps bufio.Writer, explicitly flushing on all writes.

type Flusher struct {
	w *bufio.Writer
}

// // NewFlusher creates a new Flusher from an io.Writer.

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer.

func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.Write(b)

	if err != nil {
		return -1, err
	}

	if err := foo.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

func handler(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
}
