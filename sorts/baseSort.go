package sorts

func BubbleSort(arr []int){
	n := len(arr)
	flag := false // 提前退出的标志位
	for i := 0; i < n; i++{
		for j := 0; j < n - i -1; j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 如果排好序直接退出j的循环，减少比较的次数
func InsertSort(arr []int){
	n := len(arr)

	for i := 1; i < n; i++{
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value{
				arr[j + 1] = arr[j] // 往后移动一位
			} else{
				break
			}
		}
		arr[j+1] = value // 插入数据（j-1）
	}
}

func InsertSortOrigin(arr []int){
	n := len(arr)

	for i := 1; i < n; i++{
		for j := 0; j < i; j++{
			if arr[j] > arr[i]{
				arr[j],arr[i] = arr[i],arr[j]
			}
		}
	}
}
