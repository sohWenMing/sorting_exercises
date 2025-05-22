package sorting_solutions

import (
	"fmt"
	"slices"
	"testing"

	listgeneration "github.co/sohWenMing/algorithm_exercises/list_generation"
)

func TestBubbleSort(t *testing.T) {
	unsortedList, copyList := generateLists(t)
	slices.Sort(copyList)
	bubbleSorted := BubbleSort(unsortedList)
	if len(bubbleSorted) != len(copyList) {
		t.Errorf("got length: %d, want length: %d", len(bubbleSorted), len(copyList))
	}
	for i, number := range bubbleSorted {
		if number != copyList[i] {
			t.Errorf("got %v\nwant %v", bubbleSorted, copyList)
		}
	}
	fmt.Println("sorted copy", copyList)
	fmt.Println("bubblesorted list", bubbleSorted)

}

func TestInsertionSort(t *testing.T) {
	unsortedList, copyList := generateLists(t)
	slices.Sort(copyList)
	bubbleSorted := InsertionSort(unsortedList)
	if len(bubbleSorted) != len(copyList) {
		t.Errorf("got length: %d, want length: %d", len(bubbleSorted), len(copyList))
	}
	for i, number := range bubbleSorted {
		if number != copyList[i] {
			t.Errorf("got %v\nwant %v", bubbleSorted, copyList)
		}
	}
	fmt.Println("sorted copy", copyList)
	fmt.Println("insertion sorted list", bubbleSorted)

}
func TestMergeSort(t *testing.T) {
	unsortedList, copyList := generateLists(t)
	slices.Sort(copyList)
	bubbleSorted := MergeSort(unsortedList)
	if len(bubbleSorted) != len(copyList) {
		t.Errorf("got length: %d, want length: %d", len(bubbleSorted), len(copyList))
	}
	for i, number := range bubbleSorted {
		if number != copyList[i] {
			t.Errorf("got %v\nwant %v", bubbleSorted, copyList)
		}
	}
	fmt.Println("sorted copy", copyList)
	fmt.Println("Merge sorted list", bubbleSorted)
}
func TestQuickSort(t *testing.T) {
	unsortedList, copyList := generateLists(t)
	slices.Sort(copyList)
	bubbleSorted := QuickSort(unsortedList)
	if len(bubbleSorted) != len(copyList) {
		t.Errorf("got length: %d, want length: %d", len(bubbleSorted), len(copyList))
	}
	for i, number := range bubbleSorted {
		if number != copyList[i] {
			t.Errorf("got %v\nwant %v", bubbleSorted, copyList)
		}
	}
	fmt.Println("sorted copy", copyList)
	fmt.Println("Quick sorted list", bubbleSorted)
}

func generateLists(t *testing.T) ([]int, []int) {
	unsortedList, err := listgeneration.GenerateList()(100, 10)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	copyList := make([]int, 0, len(unsortedList))
	for _, number := range unsortedList {
		copyList = append(copyList, number)
	}
	return unsortedList, copyList
}
