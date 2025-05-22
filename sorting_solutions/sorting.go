package sorting_solutions

// ########################## Bubble Sort Solution Start ################################ //
func BubbleSort(input []int) []int {
	lastIdx := len(input) - 1
	for {
		swapped := false
		// for every loop, restart i at 0
		for i := 0; i < lastIdx; i++ {
			if input[i+1] < input[i] {
				input[i], input[i+1] = input[i+1], input[i]
			}
			swapped = true
		}
		if swapped == false {
			break
		}
		lastIdx--
	}
	return input
}

// ########################## Bubble Sort Solution End ################################ //

// ########################## Insertion Sort Solution Start ################################ //
func InsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		curr := i
		for curr > 0 && input[curr] < input[curr-1] {
			input[curr], input[curr-1] = input[curr-1], input[curr]
			curr--
		}
	}
	return input
}

// ########################## Insertion Sort Solution End ################################ //

// ########################## Merge Sort Solution Start ################################ //
func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	middle := len(input) / 2
	return merge(MergeSort(input[:middle]), MergeSort(input[middle:]))
}

func merge(first, second []int) []int {
	returned := make([]int, 0, len(first)+len(second))
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			returned = append(returned, first[i])
			i++
		} else {
			returned = append(returned, second[j])
			j++
		}
	}
	for i < len(first) {
		returned = append(returned, first[i])
		i++
	}
	for j < len(second) {
		returned = append(returned, second[j])
		j++
	}
	return returned
}

// ########################## Merge Sort Solution End ################################ //

// ########################## Quick Sort Solution Start ################################ //
func QuickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	quicksort_internal(input, 0, len(input)-1)
	return input
}

func quicksort_internal(input []int, low, high int) {
	if low >= high {
		return
	}
	middle := getPivot(input, low, high)
	quicksort_internal(input, low, middle-1)
	quicksort_internal(input, middle+1, high)
	return
}

func getPivot(input []int, low, high int) int {
	pivot := input[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if input[j] < pivot {
			i++
			input[i], input[j] = input[j], input[i]
		}
	}
	middle := i + 1
	input[middle], input[high] = input[high], input[middle]
	return middle
}

// ########################## Quick Sort Solution End ################################ //

// ########################## Selection Sort Solution Start ################################ //
func SelectionSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		smallest_idx := i
		for j := smallest_idx + 1; j < len(input); j++ {
			if input[j] < input[smallest_idx] {
				smallest_idx = j
			}
		}
		input[i], input[smallest_idx] = input[smallest_idx], input[i]
	}
	return input
}

// ########################## Selection Sort Solution End ################################ //
