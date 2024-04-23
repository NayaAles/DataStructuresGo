package Array

import (
	"math/rand"
)

func CreateArray(count int) []int {
	newArray := make([]int, count)

	for i := 0; i < count; i++ {
		random := rand.Intn(100)
		newArray[i] = random
	}

	return newArray
}

func AppendItem(array []int, item int) []int {
	newArray := make([]int, len(array)+1)

	for i := 0; i < len(array); i++ {
		newArray[i] = array[i]
	}

	newArray[len(array)] = item

	return newArray
}

func InsertItem(array []int, item int, index int) []int {
	newArray := make([]int, len(array)+1)

	for i := 0; i < index; i++ {
		newArray[i] = array[i]
	}

	newArray[index] = item

	for i := index + 1; i < len(array)+1; i++ {
		newArray[i] = array[i-1]
	}

	return newArray
}

func RemoveItemInArray(array []int, index int) []int {
	newArray := make([]int, len(array)-1)

	for i := 0; i < index; i++ {
		newArray[i] = array[i]
	}

	for i := index + 1; i < len(array); i++ {
		newArray[i-1] = array[i]
	}

	return newArray
}

func SelectItem(array []int, item int) int {
	// Пузырьковая сортировка O(n^2)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	// Бинарный поиск O(log n)
	start := 0
	end := len(array) - 1

	for start <= end {
		mid := start + (end-start)/2

		if array[mid] < item {
			start = mid + 1
		} else if array[mid] > item {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func Max(array []int) int {
	// Пузырьковая сортировка O(n^2)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array[len(array)-1]
}

func Min(array []int) int {
	// Пузырьковая сортировка в обратную сторону O(n^2)
	for i := len(array); i > 0; i-- {
		for j := len(array); j > i+1; j-- {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array[0]
}
