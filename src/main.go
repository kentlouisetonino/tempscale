package main

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func main() {
	displays.ClearScreen()
	displays.AddNewline()
	fmt.Println("\t", displays.ColorBlue, "---------------------------------------------------------", displays.ColorReset)
	displays.AddNewline()
	fmt.Println("\t\t\t\t", displays.ColorGreen, "TempScale", displays.ColorReset)
	displays.AddNewline()
	fmt.Println("\t\t", "A CLI tool for converting temperature scale.")
	displays.AddNewline()
	fmt.Println("\t", displays.ColorBlue, "---------------------------------------------------------", displays.ColorReset)
	displays.AddNewline()
}
