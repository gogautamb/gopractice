package main

import (
	"fmt";
	"math"
) 

func setMinMax(min *int, max *int, smaller int, larger int) {
	if (*min > smaller) {
		*min = smaller
	}

	if (*max < larger) {
		*max = larger
	}
}

func findMinMax(a []int) (min int, max int) {
	min, max = math.MaxInt, math.MinInt

	size := len(a)

	i, j := 0, 1

	for k := 0; k < size/2; k++ {
		if a[i] < a[j] {
			setMinMax(&min, &max, a[i], a[j]) 
		} else {
			setMinMax(&min, &max, a[j], a[i])
		}
		i += 2
		j += 2
	}

	if (size%2 != 0) {
		setMinMax(&min, &max, a[size - 1], a[size - 1])
	}
	return min, max
}

func main () {

	array := []int{ 12, 34, -13, 450, 89, 2, 0, 13, 10, 99, -100, -222, -552, 300 }

	var min, max int

	min, max = findMinMax(array)

	fmt.Printf("Min: %d and Max: %d\n", min, max)

}