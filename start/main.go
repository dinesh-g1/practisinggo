package main

import (
	"fmt"
	"github.com/gdinesh/banner"
	p "github.com/gdinesh/palindrome"
)

func main() {
	runes and bytes in strings
	s := "తెలుగులో 56 అక్షరాలు ఏవి"
	lenS := 0
	for _ , _ = range s {
		//fmt.Print(string(r), ":")
		lenS++
	}
	banner.Banner(s, lenS + 10)


	//Palindrome
	fmt.Println(p.IsPalindrome("A man, a plan, a canal: Panama"))
}