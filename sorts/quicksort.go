package sorts

func QuickSort(arr []int){
	separateSort(arr,0,len(arr)-1)
}

func separateSort(arr []int,start,end int){
	if start >= end{
		return
	}
	//partial := Partion(arr,start,end)
	partial := PartionUseStartIndex(arr,start,end)
	separateSort(arr,start,partial-1)
	separateSort(arr,partial+1,end)
}

func Partion(arr []int, start,end int) int{
	// 选取最后一个数字当做参考数字
	privot := arr[end]

	i := start
	for j := start; j < end;j++{
		if arr[j] < privot{
			if !(i == j){
				// 交换位置
				arr[i],arr[j] = arr[j],arr[i]
			}
			i++
		}
	}

	arr[i],arr[end] = arr[end],arr[i]
	return i
}

func PartionUseStartIndex(arr []int, start,end int) int{
	privot := arr[start]
	i := start
	j := end

	for{
		if (i < j){
			for{
				if i < j && privot <= arr[j]{
					j --
				}else{
					arr[i] = arr[j]
					break
				}
			}
			for{
				if i < j && arr[i] <= privot{
					i++
				}else{
					arr[j] = arr[i]
					break
				}
			}
		} else{
			break
		}
	}
	arr[i] = privot
	return i
}