package main

import "fmt"

func main() {
	var name string = "hello"
	var age int = 25
	var score float64 = 98.6
	var active bool = true

	var city = "Delhi"
	var population = 600000

	language := "Go"
	version := 1.26

	var result string
	result = "computed later"

	var (
		host = "localhost"
		port int = 8000
		debug bool = false
	)

	fmt.Println(name, age, score, active)
	fmt.Println(city, population)
	fmt.Println(language, version)
	fmt.Println(host, port, debug)
	fmt.Println(result)
}