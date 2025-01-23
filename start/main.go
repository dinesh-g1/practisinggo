package main

import (
	//"fmt"
	"github.com/gdinesh/banner"
)

func main() {
	s := "తెలుగులో 56 అక్షరాలు ఏవి"
	lenS := 0
	for _ , _ = range s {
		//fmt.Print(string(r), ":")
		lenS++
	}
	banner.Banner(s, lenS + 10)
}