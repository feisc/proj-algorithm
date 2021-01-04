package sorts

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{5,7,4,3,8,2}
	MergeSort(arr)
	t.Log(arr)
}