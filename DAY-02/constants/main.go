package main

import "fmt"

const Pi = 3.14

const (
	StatusOK    = 200
	Status404   = 404
	StatusError = 500
)

// iota is an auto-incrementing counter (starts at 0, +1 per line)
const (
	_  = iota          // iota = 0, discarded via _
	KB = 1 << (10 * iota) // iota = 1 → 1 << 10 = 1024
	MB                     // iota = 2 → 1 << 20 = 1,048,576 (repeats prev pattern)
	GB                     // iota = 3 → 1 << 30 = 1,073,741,824
)

func main() {
	const greeting = "Hello"

	fmt.Println(Pi)
	fmt.Println(greeting)

	const typed int = 42
	const untyped = 42
	fmt.Printf("typed: %T, untyped: %T\n", typed, untyped)

	fmt.Println(StatusOK, Status404, StatusError)

	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
}
