package InsertionSort

//confusing but ok i guess
//less steps then bubble sort
//speed: O(n^2)
//best case can be O(n) compared to selection sort's O(n^2)
//small data sets: ok
//big data sets: *skull emoji*
//preferable to both bubble sort and selection sort

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}

	return arr
}
