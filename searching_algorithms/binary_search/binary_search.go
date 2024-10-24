package BinarySearch

//amazing at bigger data sets
//slower then linear on small data sets
//speed: O(log(n))

func binarySearch(arr []int, target int) int {
	var middle int
	var hi int = len(arr) - 1
	var lo int = 0

	for lo <= hi {
		middle = (lo + hi) / 2
		if target == arr[middle] {
			return middle
		} else if target < arr[middle] {
			hi = middle - 1
		} else {
			lo = middle + 1
		}
	}

	return -1
}
