package displays

import "fmt"

func AppOptionInputError() {
	fmt.Println(Tab, ColorRed, " Invalid input. Please try again.", ColorReset)
}
