package sorts

import "math/rand"

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := MergeSort(arr[:len(arr)/2])
	second := MergeSort(arr[len(arr)/2:])
	return Merge(first, second)
}

func Merge(arr []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(arr) && j < len(b) {
		if arr[i] < b[j] {
			final = append(final, arr[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(arr); i++ {
		final = append(final, arr[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
