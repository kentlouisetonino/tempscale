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

	// Display the options.
	displays.AppOptions()
	displays.AddNewline()

	// Choose a temperature scale to convert.
	displays.AppMainOption()
}
