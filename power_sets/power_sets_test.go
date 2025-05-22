package powersets

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestPowerSets(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected [][]int
	}

	tests := []test{
		{
			"basic test",
			[]int{1, 2, 3},
			[][]int{
				{},
				{1},
				{1, 2},
				{1, 3},
				{1, 2, 3},
				{2},
				{2, 3},
				{3},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getSubStringStringMap(GetPowerSets(test.input))
			want := getSubStringStringMap(test.expected)

			if len(got) != len(want) {
				fmt.Printf("got %v\n want %v", got, want)
				t.Errorf("\nlength of power sets don't match: got %d, want %d", len(got), len(want))
			}
			for key, _ := range got {
				if _, ok := want[key]; !ok {
					t.Errorf("power sets do not match\n got %v want %v", got, want)
				}
			}
		})
	}
}

func makeStringFromSubset(subset []int) string {
	sort.Ints(subset)
	stringSlice := make([]string, len(subset))
	for _, idxVal := range subset {
		stringSlice = append(stringSlice, fmt.Sprintf("%d", idxVal))
	}
	return strings.Join(stringSlice, "")

}

func getSubStringStringMap(pow_set [][]int) map[string]struct{} {
	returnedMap := make(map[string]struct{})
	for _, subset := range pow_set {
		key := makeStringFromSubset(subset)
		returnedMap[key] = struct{}{}
	}
	return returnedMap
}
