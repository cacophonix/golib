package stringutil

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func z_index(arr []int) []int {
	n := len(arr)
	ret := make([]int, n)

	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			ret[i] = min(r-i+1, ret[i-l])
		}
		for i+ret[i] < n && arr[ret[i]] == arr[i+ret[i]] {
			ret[i]++
		}
		if i+ret[i]-1 > r {
			l = i
			r = i + ret[i] - 1
		}
	}

	return ret
}
