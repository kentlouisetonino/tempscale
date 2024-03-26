package displays

import "fmt"

func AppDescription() {
	fmt.Println("\t", ColorBlue, "---------------------------------------------------------", ColorReset)
	AddNewline()
	fmt.Println("\t\t\t\t", ColorGreen, "TempScale", ColorReset)
	AddNewline()
	fmt.Println("\t\t", "A CLI tool for converting temperature scale.")
	AddNewline()
	fmt.Println("\t", ColorBlue, "---------------------------------------------------------", ColorReset)
}
