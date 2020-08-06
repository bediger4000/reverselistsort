package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var lst []int
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err == nil {
			lst = append(lst, n)
		}
	}

	fmt.Printf("List:        %v\n", lst)
	sort(lst)
	fmt.Printf("Sorted list: %v\n", lst)
}

func sort(ary []int) {

	length := len(ary)

	for n := length - 1; n != 0; n-- {
		for i := 0; i < n; i++ {
			if ary[i] < ary[i+1] {
				reverse(ary, i, i+1)
			}
		}
	}
}
