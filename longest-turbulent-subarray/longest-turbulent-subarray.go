package longest_turbulent_subarray

func maxTurbulenceSize(arr []int) int {
	ans := 1
	anchor := 0
	for i := 1; i < len(arr); i++ {
		c := compare(arr[i-1], arr[i])
		if c == 0 {
			anchor = i
		} else if i == len(arr)-1 || c*compare(arr[i], arr[i+1]) != -1 { // if the next array element is not flipped
			ans = max(ans, i-anchor+1)
			anchor = i
		}
	}

	return ans
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func compare(a int, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
