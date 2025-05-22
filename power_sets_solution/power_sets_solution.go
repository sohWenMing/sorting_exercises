package powersets_solution

import "fmt"

func GetPowerSets(input []int) [][]int {
	if len(input) == 0 {
		return [][]int{{}}
	}
	returned := [][]int{}
	cur_chunk := []int{input[0]}
	subset := makeCopySlice(input[1:])
	fmt.Println("subset:", subset)
	returned_from_func := GetPowerSets(subset)
	for _, ret_set := range returned_from_func {
		to_add := append(cur_chunk, ret_set...)
		returned = append(returned, to_add)
		returned = append(returned, ret_set)
	}
	return returned
}

func makeCopySlice(input []int) []int {
	returned := make([]int, len(input))
	copy(returned, input)
	return returned
}
