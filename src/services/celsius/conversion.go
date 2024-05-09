package celsius

import "github.com/kentlouisetonino/tempscale/src/displays"

func Conversion() {
	// Clear the screen.
	displays.ClearScreen()
	displays.AddNewline()
	displays.AddNewline()

	// Show the description.
	displays.CelsiusOptionDescription()
	displays.AddNewline()
	displays.AddNewline()
}
