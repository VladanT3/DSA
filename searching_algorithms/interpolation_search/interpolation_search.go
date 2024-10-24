package InterpolationSearch

//fancy binary search
//works really good with uniformally distributed data (e.g. [1, 2, 3, 4, 5, 6, 7, 8])
//slower on not uniformally distributed data (e.g. [1, 2, 4, 8, 16, 32, 64, 128, 256, 512])
//average case: O(log(log(n)))
//worst case: O(n) -> if values in data set increased exponentially

func interpolationSearch(arr []int, target int) int {
	var hi int = len(arr) - 1
	var lo int = 0

	for target >= arr[lo] && target <= arr[hi] && lo <= hi {
		var probe int = lo + (hi-lo)*(target-arr[lo])/(arr[hi]-arr[lo])

		if arr[probe] == target {
			return probe
		} else if arr[probe] < target {
			lo = probe + 1
		} else {
			hi = probe - 1
		}
	}

	return -1
}
