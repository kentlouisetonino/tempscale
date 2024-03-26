package displays

import "fmt"

func AppDescription() {
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
	AddNewline()
	fmt.Println(Tab, Tab, Tab, Tab, ColorGreen, "TempScale", ColorReset)
	AddNewline()
	fmt.Println(Tab, Tab, "A CLI tool for converting temperature scale.")
	AddNewline()
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
}
