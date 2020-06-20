package main

import (
	"fmt"

	"github.com/armarom/finalexam/todo"
)

func main() {
	fmt.Println("Go Programming Language Final exam...\n")
	r := todo.SetupRouter()
	r.Run(":2019")

}
