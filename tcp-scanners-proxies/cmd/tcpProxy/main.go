package main

// Building a TCP Proxy
// First let's start with a echo server

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
