package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{5,7,3,6,8,4}
	BubbleSort(arr)
	t.Log(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{5,7,3,6,8,4}
	InsertSort(arr)
	t.Log(arr)
}

func TestInsertSortOrigin(t *testing.T) {
	arr := []int{5,7,3,6,8,4}
	InsertSortOrigin(arr)
	t.Log(arr)
}