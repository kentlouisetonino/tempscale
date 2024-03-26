package main

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func main() {
	displays.DisplayNewline()
	fmt.Println("\t--------------------------------------------------")
	displays.DisplayNewline()
	fmt.Println("                          TempScale                 ")
	displays.DisplayNewline()
	fmt.Println("\t--------------------------------------------------")
}
