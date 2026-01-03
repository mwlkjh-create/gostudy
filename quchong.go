package main

import "fmt"

func quchong(a []int) []int {
	s := make(map[int]bool)
	result := []int{}
	for _, v := range a {
		if s[v] != true {
			s[v] = true
			result = append(result, v)
		}
	}
	return result
}
func main() {
	b := []int{1, 5, 3, 3, 5, 6, 10, 8, 9, 10}
	fmt.Println(quchong(b))
}
