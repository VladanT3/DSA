package QuickSort

//move everything smaller then the pivot to the left of it and bigger to the right of it, then repeat
//speed: O(n*log(n)) or O(n^2) if already sorted or very close to
//space: O(log(n))

func quickSort(arr []int, start, end int) {
	if end <= start { //base case
		return
	}

	var pivot int = partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	var i, j int = start - 1, start
	var pivot int = arr[end]

	for j <= end {
		if arr[j] < pivot || j == end {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
		j++
	}

	return i
}
