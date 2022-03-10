package main

import (
	"fmt"
)

func PrintSpriralOrder(A [][]int) {
	i,m := 0,len(A) - 1

	for k := 0; k < len(A)/2 ; k++ {
		fmt.Println("---- left to right ------", k)
		for x := i; x < m; x++ {
			fmt.Printf("A[%d][%d]=%d\n", i, x, A[i][x])
		}

		fmt.Println("---- top to bottom ------", k)

		for x := i; x < m; x++ {
			fmt.Printf("A[%d][%d]=%d\n", x, m, A[x][m])
		}

		fmt.Println("---- right to left ------", k)

		for x := m; x > i; x-- {
			fmt.Printf("A[%d][%d]=%d\n", m, x, A[m][x])
		}

		fmt.Println("---- bottom to top ------", k)

		for x := m; x > i; x-- {
			fmt.Printf("A[%d][%d]=%d\n", x, i, A[x][i])
		}

		i += 1
		m -= 1		
	} 

	if (len(A)%2 != 0) {
		fmt.Println("--- only ---")
		fmt.Printf("A[%d][%d]=%d\n", i, i, A[i][i])
	}

}

func main() {
	fmt.Println("Example of 2D Slice")

	matrix := [][]int {
		{1, 2, 3, 4, 20},
		{5, 6, 7, 8, 21},
		{9, 10, 11, 12, 22},
		{13, 14, 15, 16, 23},
		{31, 32, 33, 34, 35},
	}

	PrintSpriralOrder(matrix)
}