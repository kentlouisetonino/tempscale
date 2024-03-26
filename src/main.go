package main

import (
	"github.com/kentlouisetonino/tempscale/src/displays"
)

func main() {
	// Clear the terminal first.
	displays.ClearScreen()
	displays.AddNewline()

	// Display the application description.
	displays.AppDescription()
	displays.AddNewline()
}
