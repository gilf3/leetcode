package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func merge(A []int, m int, B []int, n int) {
	i, j, h := m-1, n-1, len(A)
	for h > 0 {
		h--
		if j < 0 {
			return
		}
		if i < 0 {
			copy(A, B[:j+1])
			return
		}
		if A[i] < B[j] {
			A[h] = B[j]
			j--
		} else {
			A[h] = A[i]
			i--
		}
	}
}
