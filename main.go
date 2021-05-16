package main

import (
	"algorithm/action"
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	item int
	next *Node
}

func createLink(li []int) Node {
	head := Node{item: li[0]}
	for i := 1; i < len(li); i++ {
		node := Node{item: li[i]}
		node.next = &head
		head = node
	}
	return head
}

func isUnique(astr string) bool {
	var hash = map[byte]int{}
	for i := 0; i < len(astr); i++ {
		if _, ok := hash[astr[i]]; ok {
			return false
		}
		hash[astr[i]] = i
	}
	return true
}

func main() {
	res := action.FindMedianSortedArrays([]int{1}, []int{1, 2, 3, 9})
	fmt.Println(res)
	binarySearch(2)

	li := []int{5, 7, 4, 6, 3, 1, 2, 9, 8, -1}
	action.QuickSort(li, 0, len(li)-1)

	action.MergeArray([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	action.MergeArray([]int{1, 0}, 1, []int{1}, 1)
	action.Fib(2)
	action.LongestPalindrome("babad")
	isUnique("leetcode")
	createLink([]int{1, 2, 3, 4, 5})
	data := GenerateInt(100000, 100000)
	start := makeTimestamp()
	// fmt.Printf("%v\n", data)
	data = action.MergeSort(data)
	fmt.Printf("cost %v ms \n", makeTimestamp()-start)

	//insertSort(array)
	//selectSort(array)

}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func binarySearch(val int) int {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	left := 0
	right := len(array) - 1

	for left <= right {
		mind := (left + right) / 2
		mindValue := array[mind]
		if mindValue == val {
			return array[mind]
		}
		if val > mindValue {
			left = mind + 1
		} else {
			right = mind - 1
		}
	}
	return 0
}

func selectSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		minLocal := i
		for j := i + 1; j < len(array); j++ {
			array_j := array[j]
			array_i := array[minLocal]
			if array_j < array_i {
				array[j], array[minLocal] = array[minLocal], array[j]
			}
		}
		println(array)
	}
	return array
}

// GenerateInt : 随机生成 n 个 最大为 max 的 Int 切片
func GenerateInt(n, max int) (data []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(max))
	}
	return
}
