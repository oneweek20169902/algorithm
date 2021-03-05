package action

import "fmt"

//斐波那契数
func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fib(n-1) + Fib(n+2)
	} else {
		return -1
	}
}

//合并排序的数组
func MergeArray(A []int, m int, B []int, n int) {
	resutl := make([]int, 0)
	var i, j int = 0, 0
	length := m + n - 1
	for i < m && j < n {
		if len(resutl) == length {
			break
		}
		if n == 0 {
			return
		} else if A[i] < B[j] {
			resutl = append(resutl, A[i])
			i++
		} else {
			resutl = append(resutl, B[j])
			j++
		}
	}
	if j < n {
		resutl = append(resutl, B[j:]...)
	} else if i < m {
		resutl = append(resutl, A[i:m]...)
	}
	A = append(A[:0], resutl...)
	fmt.Println(A)
}
