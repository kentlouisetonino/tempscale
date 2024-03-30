package main

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func main() {
	mainOption := ""

	// Clear the terminal first.
	displays.ClearScreen()
	displays.AddNewline()
	displays.AddNewline()

	// Display the application description.
	displays.AppDescription()
	displays.AddNewline()

	// Display the options.
	displays.AppOptions()
	displays.AddNewline()

	// Choose a temperature scale to convert.
	mainOption = displays.AppMainOption()
	displays.AddNewline()
	displays.AddNewline()
	displays.AddNewline()

	// Handle the input.
	if mainOption == "-1" {
		fmt.Println("Invalid Input.")
	} else {
		fmt.Println(mainOption)
	}
}
