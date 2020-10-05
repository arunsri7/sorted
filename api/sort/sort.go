package sort

import "sort"

//implements BubbleSort
func BubbleSort(a []int) [][]int {
	//2-D temp array to store the sorted array after each iteration
	var sorted [][]int
	for j := 0; j < len(a); j++ {
		var b []int
		count := 0
		for i := range a {
			if i == len(a)-1 {
				continue
			}
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				count++
			}
		}
		//To avoid looping over the sorted array
		if count == 0 {
			return sorted
		}
		//Appending the value of 'a' to 'sorted'
		b = make([]int, len(a))
		copy(b, a)
		sorted = append(sorted, b)
	}
	return sorted
}

//implements Selection sort
func SelectionSort(unSorted []int) [][]int {
	//2-D temp array to store the sorted array after each iteration
	var b []int
	var sorted [][]int
	var n = len(unSorted)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if unSorted[j] < unSorted[minIdx] {
				minIdx = j
			}
		}
		unSorted[i], unSorted[minIdx] = unSorted[minIdx], unSorted[i]
		//Appending the value of 'a' to 'sorted'
		b = make([]int, len(unSorted))
		copy(b, unSorted)
		sorted = append(sorted, b)
		if sort.IntsAreSorted(unSorted) {
			return sorted
		}
	}
	return sorted
}
