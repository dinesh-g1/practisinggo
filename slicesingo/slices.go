package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int //int slice
	fmt.Println(len(s))

	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{2, 59, 8, 6, 3, 18, 45} // Declare and Initialize
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation
	fmt.Printf("s3 = %#v\n", s3)

	// s4 := s2[:100] // panic happens

	s3 = append(s3, 64)
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2)

	s4 := []int{}
	for i := 0; i < 10; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Printf("%#v\n", s4)

	s5 := make([]int, 5)
	for i := 0; i < len(s5); i++ {
		fmt.Printf("s5[%d] = %d\n", i, s5[i])
	}
	fmt.Printf("%#v\n", s5)

	//concating the string using slices
	fmt.Println(concat([]string{"hanuman", "lakshmana", "rama"}, []string{"sita"}))

	//sorting
	fmt.Println(median([]float64{1,5,64,3,6,36,56,34}))
}

func median(values []float64) float64 {
	sort.Float64s(values)
	i := len(values)/2

	if len(values)%2 == 1 {
		return values[i]
	}

	return (values[i] + values[i+1]) / 2
}

func appendInt(s []int, v int) []int {
	i := len(s)

	if len(s) >= cap(s) {
		//fmt.Printf("oldlen(s)=%d, oldcap(s)=%d\n", len(s), cap(s))
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		//fmt.Printf("newlen(s2)=%d, newcap(s2)=%d\n", len(s2), cap(s2))
		s = s2[:len(s)+1]
	} else {
		s = s[:len(s)+1]
	}

	s[i] = v
	//fmt.Printf("%#v\n", s)

	return s
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s;
}
