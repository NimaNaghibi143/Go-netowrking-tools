package main

import (
	"fmt"
	"sync"
)

// step #1
// a basic port scanner

// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)

// 		conn, err := net.Dial("tcp", address)

// 		if err != nil {
// 			// port is closed or filtered
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%d open:\n", i)
// 	}
// }

// step #2
// The “Too Fast” Scanner Version using goroutines
// the threads are way to faster than the connection stablishements and the funcs exit when the packets
// are still in the flight.
// one way to solve this is using the WaitGroup from the sync package which is a thread safe way to control concurrency.

// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		go func(j int) {
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("Some shit went wrong")
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%d open\n", j)
// 		}(i)
// 	}
// }

// step #3
// Synchronized Scanning Using WaitGroup

// func main() {
// 	// synchronized counter
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 1024; i++ {
// 		// the counter gets incremented each time a goroutine is created to scan a port
// 		wg.Add(1)
// 		go func(j int) {

// 			// Decrement the counter when ever a thread has finished the job
// 			defer wg.Done()
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)

// 			if err != nil {
// 				fmt.Printf("Some shit went wrong!")
// 				return
// 			}

// 			conn.Close()
// 			fmt.Printf("%d open\n", j)
// 		}(i)
// 	}

// 	// blocks unitl the counter is zero
// 	wg.Wait()
// }

// step #4
// Port Scanning Using a Worker Pool
// To avoid inconsistencies, you’ll use a pool of goroutines to manage the
// concurrent work being performed.

// The channel will be used
// to receive work, and the WaitGroup will be used to track when a single work
// item has been completed.

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}
