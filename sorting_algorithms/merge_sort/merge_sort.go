package MergeSort

//divide and conquer algorithm, uses recursion
//speed: O(n*log(n))
//faster on large data sets then bubble, selection and insertion sort
//but uses more memory then them: O(n) compared to O(1); because of calling more functions and creating subarrays

func mergeSort(arr []int) {
	if len(arr) <= 1 { //base case
		return
	}

	var middle int = len(arr) / 2
	var left_arr, right_arr []int
	var j int = 0

	for i := range arr {
		if i < middle {
			left_arr[i] = arr[i]
		} else {
			right_arr[j] = arr[i]
			j++
		}
	}

	mergeSort(left_arr)
	mergeSort(right_arr)
	merge(left_arr, right_arr, arr)
}

func merge(left_arr, right_arr, arr []int) {
	var i, l, r int = 0, 0, 0

	for l < len(left_arr) && r < len(right_arr) {
		if left_arr[l] < right_arr[r] {
			arr[i] = left_arr[l]
			i++
			l++
		} else {
			arr[i] = right_arr[r]
			i++
			r++
		}
	}

	//if some elements are left alone and the other array is "empty" and cant be compared to anything
	for l < len(left_arr) {
		arr[i] = left_arr[l]
		i++
		l++
	}
	for r < len(right_arr) {
		arr[i] = right_arr[r]
		i++
		r++
	}
}
