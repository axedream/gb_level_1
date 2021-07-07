package sorting

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}


func TestSort(t *testing.T)  {
	var data = []int{50,6,8,12,56,34,5}
	got := SortIntSlice(data)
	want := []int{5,6,8,12,34,50,56}

	if !Equal(want,got) {
		t.Errorf("Got: %d Want: %d",got,want)
	}
	return
}