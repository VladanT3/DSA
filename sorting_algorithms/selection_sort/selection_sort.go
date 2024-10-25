package SelectionSort

//go through array find smallest number and place it at the beginning, then start from next position, find smallest and replace, and so on...
//speed: O(n^2)
//small data sets - ok
//big data sets - no no

func selectionSort(arr []int) []int {
	var minimum int
	for i := 0; i < len(arr)-1; i++ {
		minimum = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minimum] {
				minimum = j
			}
		}
		var temp int = arr[i]
		arr[i] = arr[minimum]
		arr[minimum] = temp
	}

	return arr
}
