package main

import (
	"fmt"
	"my-gram/config"
)

func main() {
	config.ReadConfig()

	fmt.Println("Hello, World!")
}
