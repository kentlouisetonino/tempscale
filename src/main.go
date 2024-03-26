package main

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func main() {
	displays.AddNewline()
	fmt.Println(displays.Tab, "--------------------------------------------------")
	displays.AddNewline()
	fmt.Println("                             TempScale                 ")
	displays.AddNewline()
	fmt.Println(displays.Tab, "   A CLI tool for converting temperature scale.   ")
	displays.AddNewline()
	fmt.Println(displays.Tab, "--------------------------------------------------")
	displays.AddNewline()
}
