package segmenttree

var max [444444]int
var min [444444]int
var arr []int

func initt(pos, l, r int) {
	if l == r {
		max[pos] = l
		min[pos] = l
		return
	}
	mid := (l + r) / 2
	initt(pos*2, l, mid)
	initt(pos*2+1, mid+1, r)

	if arr[max[pos*2]] > arr[max[pos*2+1]] {
		max[pos] = max[pos*2]
	} else {
		max[pos] = max[pos*2+1]
	}

	if arr[min[pos*2]] < arr[min[pos*2+1]] {
		min[pos] = min[pos*2]
	} else {
		min[pos] = min[pos*2+1]
	}
}

func findMax(pos, l, r, ll, rr int) int {
	if l >= ll && r <= rr {
		return max[pos]
	}

	mid := (l + r) / 2
	if mid < ll {
		return findMax(pos*2+1, mid+1, r, ll, rr)
	}
	if rr <= mid {
		return findMax(pos*2, l, mid, ll, rr)
	}

	r1 := findMax(pos*2, l, mid, ll, rr)
	r2 := findMax(pos*2+1, mid+1, r, ll, rr)

	if arr[r1] > arr[r2] {
		return r1
	}
	return r2

}

func findMin(pos, l, r, ll, rr int) int {
	if l >= ll && r <= rr {
		return min[pos]
	}

	mid := (l + r) / 2
	if mid < ll {
		return findMin(pos*2+1, mid+1, r, ll, rr)
	}
	if rr <= mid {
		return findMin(pos*2, l, mid, ll, rr)
	}

	r1 := findMin(pos*2, l, mid, ll, rr)
	r2 := findMin(pos*2+1, mid+1, r, ll, rr)

	if arr[r1] < arr[r2] {
		return r1
	}
	return r2

}
