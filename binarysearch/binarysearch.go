package binarysearch

func BinarySearch(a []int, v int) int{
	// 默认所传的数是正整数，如果有负数，返回一个负数的最大值
	length := len(a)
	if length == 0{
		return -1
	}

	low := 0
	high := length -1
	for low <= high{
		mid := low + ((high - low) >> 1)
		if a[mid] == v{
			return mid
		}else if a[mid] > v{
			high = mid - 1
		}else{
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个给定定值的元素
func BinarySearchFirst(a []int, v int) int{
	length := len(a)
	if length == 0{
		return -1
	}
	low := 0
	high := length - 1
	for low <= high{
		mid := low + ((high - low) >> 1)
		if a[mid] > v{
			high = mid - 1
		}else if a[mid] < v{
			low = mid + 1
		}else{
			if mid == 0 || a[mid - 1] != v{
				return mid
			}else{
				high = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个给定的数
func BinarySearchLast(a []int, v int) int {
	length := len(a)
	if length == 0{
		return  -1
	}
	low := 0
	high := length - 1
	for low <= high{
		mid := low +((high - low) >> 1)
		if a[mid] > v{
			high = mid - 1
		}else if a[mid] < v{
			low = mid + 1
		}else{
			if mid == length - 1 || a[mid + 1] != v {
				return mid
			}else{
				low = mid + 1
			}
		}
	}
	return - 1
}

// 第一个比指定元素大
func BinarySearchFirstGT(a []int, v int) int{
	length := len(a)
	if length == 0{
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] > v {
			high = mid -1
		}else if a[mid] < v {
			low = mid + 1
		}else{
			// 如果查找的是最后一个，直接返回
			if mid == length - 1{
				return mid
			}
			if a[mid + 1] > v{
				return mid + 1
			}else{
				low = mid + 1
			}
		}
	}
	return -1
}

// 最后一个比指定的元素要小
func BinarySearchLastLT(a []int, v int) int {
	length := len(a)
	if length == 0{
		return -1
	}
	low := 0
	high := length - 1
	for low <= length{
		mid := low + ((high - low) >> 1)
		if a[mid] > v{
			high = mid - 1
		}else if a[mid] < v{
			low = mid + 1
		}else{
			if mid == 0{
				return mid
			}
			if a[mid - 1] < v{
				return mid - 1
			}else{
				high = mid - 1
			}
		}
	}
	return -1
}

