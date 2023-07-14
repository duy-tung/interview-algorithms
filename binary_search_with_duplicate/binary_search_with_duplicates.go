package binary_search_with_duplicates

func binary_search_with_duplicates(arr []int, target int) int {
	min, max := 0, len(arr)-1
	result := -1

	for max >= min {
		mid := min + (max-min)/2
		if arr[mid] == target {
			max = mid - 1
			result = mid
		} else if arr[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return result
}
