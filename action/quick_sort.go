package action

func QuickSort(array []int, left int, right int) {
	if left < right {
		mid := partition(array, left, right)
		QuickSort(array, left, mid-1)
		QuickSort(array, mid+1, right)
	}
}

//归位函数算出数组中间的的下标位置
func partition(array []int, left int, right int) int {
	temp := array[left]
	for left < right {
		for left < right && array[right] >= temp {
			right -= 1
		}
		array[left] = array[right]
		for left < right && array[left] <= temp {
			left += 1
		}
		array[right] = array[left]
		array[left] = temp
	}
	return left
}
