package main

/*
 * reverse(lst, i, j), which reverses lst
 * from i to j, inclusive
 */

func reverse(lst []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		lst[i], lst[j] = lst[j], lst[i]
	}
}
