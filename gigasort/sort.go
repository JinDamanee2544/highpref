package main

import "sort"

func QuickSort(arr []KeyPair, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func partition(arr []KeyPair, low int, high int) int {
	pivot := arr[high].Value
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j].Key < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func NativeSort(arr []KeyPair) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value < arr[j].Value
	})
}
