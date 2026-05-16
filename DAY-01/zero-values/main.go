package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var s1 []string

	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("string: %q\n", s)
	fmt.Printf("pointer: %v\n", p)
	fmt.Printf("slice: %v\n", s1)

	type Server struct {
		Host string
		Port int
		Running bool
	}

	var srv Server
	fmt.Printf("server: %+v\n", srv)
}