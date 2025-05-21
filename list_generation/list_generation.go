package listgeneration

import (
	"fmt"
	"math/rand"
)

func getRandNum(max int) int {
	return rand.Intn(max)
}

func GenerateList() func(max, length int) (list []int, err error) {
	return func(max, length int) (list []int, err error) {
		if length < 1 {
			return nil, fmt.Errorf("length of list has to be positive")
		}
		returnedList := make([]int, 0, length)
		for i := 0; i < length; i++ {
			returnedList = append(returnedList, getRandNum(max))
		}
		return returnedList, nil
	}
}
