package BubbleSort

//comparing and switching two adjecent elements
//speed: O(n^2)
//for small data sets - meh, ok
//for big data sets - hell no

func bubbleSort(arr []int) []int {
	var i, j int
	for i = len(arr) - 1; i > 0; i-- {
		for j = 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				var temp int = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}
