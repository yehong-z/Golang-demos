package util

func Swap(a []int, u int, v int) {
	t := a[u]
	a[u] = a[v]
	a[v] = t
}

func Reverse(a []int, l int, r int) {
	var t int
	for l < r {
		t = a[l]
		a[l] = a[r]
		a[r] = t
		l++
		r--
	}
}
