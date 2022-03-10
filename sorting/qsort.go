package main

import ("fmt")

func swap(A []int, a, b int) {
	if (a != b) {
		temp := A[a]
		A[a] = A[b]
		A[b] = temp
	}
}

func partition(A []int, left, right int) (pivot int) {
	i:= left - 1

	p := A[right]

	for j := left; j < right; j++ {
		if A[j] < p {
			i += 1
			swap(A, i, j)
		}
	} 

	i += 1
	swap(A, i, right)

	return i
}

func QuickSort(A []int) {
	quickSortUtil(A, 0, len(A) - 1)
}

func quickSortUtil(A []int, left, right int) {
	if (left < right) {
		piv := partition(A, left, right)
		quickSortUtil(A, left, piv - 1)
		quickSortUtil(A, piv + 1, right)
	}
}

func main() {
	A := []int {4, 6, 15, 15, 4, 4,67, -23, 12}
	fmt.Println(A)
	QuickSort(A)
	fmt.Println(A)
}