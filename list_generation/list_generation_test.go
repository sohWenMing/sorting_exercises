package listgeneration

import "testing"

func TestGenerateList(t *testing.T) {
	returnedList, err := GenerateList()(100, 10)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	if len(returnedList) != 10 {
		t.Errorf("got %d, want %d", len(returnedList), 10)
	}
}
