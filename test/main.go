package main

import (
	"fmt"

	"goreloaded"
)

func main() {
	goreloaded.Goreloaded()
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}
