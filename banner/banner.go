package banner

import (
	"fmt"
	"unicode/utf8"
)

func Banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text))/2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
}