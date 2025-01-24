package main

import (
	"fmt"

	"github.com/gdinesh/banner"
	"github.com/gdinesh/github"
	p "github.com/gdinesh/palindrome"
	"github.com/gdinesh/sha1"
)

func main() {
	//runes and bytes in strings
	s := "తెలుగులో 56 అక్షరాలు ఏవి"
	lenS := 0
	for _, _ = range s {
		//fmt.Print(string(r), ":")
		lenS++
	}
	banner.Banner(s, lenS+10)

	//Palindrome
	fmt.Println(p.IsPalindrome("A man, a plan, a canal: Panama"))

	//API Request Response explorer
	github.APIPractice()

	// Working with Reader and Writer
	fileName := "sha1/http.log.gz"
	fmt.Println(sha1.Sha1Sum(fileName))
}
