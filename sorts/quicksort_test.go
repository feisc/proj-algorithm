package sorts

import "testing"

func TestQucikSort(t *testing.T){
	arr := []int{5,7,4,3,8,2}
	QuickSort(arr)
	t.Log(arr)
}
