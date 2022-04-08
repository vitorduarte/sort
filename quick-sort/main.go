package main

import (
	"fmt"
)

const (
	Lomoto QuickSortMethod = iota
	Hoare
)

func quickSort(slice []int, method QuickSortMethod) {
	fmt.Printf("unsorted slice: %v\n", slice)

	if method == Lomoto {
		quickSortLomuto(slice, 0, len(slice)-1)
	}
	if method == Hoare {
		quickSortHoare(slice, 0, len(slice)-1)
	}

	fmt.Printf("sorted slice: %v\n\n", slice)
}

func quickSortLomuto(slice []int, lo int, hi int) {
	// Ensure indices are in correct order
	if lo >= hi || lo < 0 || hi > len(slice) {
		return
	}

	// Partition array and get the pivot index
	p := partitionLomuto(slice, lo, hi)

	// Sort the two partitions
	quickSortLomuto(slice, lo, p-1)
	quickSortLomuto(slice, p+1, hi)

}

func partitionLomuto(slice []int, lo int, hi int) int {
	pivot := slice[hi] // Choose the last element

	// Temporary pivot index
	i := lo - 1

	for j := lo; j <= hi-1; j++ {
		// If the current element is less than or equal to the pivot
		if slice[j] <= pivot {
			// Move the temporry pivot index forward
			i += 1

			// Swap the current element with the element at the temporry pivot index
			slice[j], slice[i] = slice[i], slice[j]
		}
	}

	// Move the pivot element to the correct pivot position (between the smaller and larger elements)
	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]
	return i // the pivot index
}

func quickSortHoare(slice []int, lo int, hi int) {
	if lo >= 0 && hi >= 0 && lo < hi {
		p := partitionHoare(slice, lo, hi)
		quickSortHoare(slice, lo, p)
		quickSortHoare(slice, p+1, hi)
	}
}

func partitionHoare(slice []int, lo int, hi int) int {
	pivot := slice[(hi+lo)/2]
	// Left index
	i := lo - 1

	// Right index
	j := hi + 1

	for {
		// Move the left index to the right at least once and while the element at
		// the left index is less than the pivot
		for {
			i += 1
			if slice[i] >= pivot {
				break
			}
		}

		// Move the right index to the left at least once and while the element at
		// the right index is greater than the pivot
		for {
			j -= 1
			if slice[j] <= pivot {
				break
			}
		}

		// If the indices crossed, return
		if i >= j {
			return j
		}

		// Swap the elements at the left and right indices
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type QuickSortMethod int

func main() {
	quickSort([]int{3, 5, 12, 34, 1, 3, 76, 3, 5, 1, 9, 2}, Lomoto)
	quickSort([]int{3, 5, 12, 34, 1, 3, 76, 3, 5, 1, 9, 2}, Hoare)
}
