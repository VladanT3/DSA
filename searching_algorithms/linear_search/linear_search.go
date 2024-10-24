package LinearSearch

//the classic
//speed: O(n)

func linearSearch(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}

	return -1
}
