package sort

func QuickSort(array []int) {
	quicksort(array, 0, len(array)-1)
}

func quicksort(array []int, l, r int) {
	if l >= r {
		return
	}
	j := l
	for i := l + 1; i <= r; i++ {
		if array[i] < array[l] {
			j++
			t := array[j]
			array[j] = array[i]
			array[i] = t
		}
	}
	t := array[j]
	array[j] = array[l]
	array[l] = t
	quicksort(array, l, j-1)
	quicksort(array, j+1, r)
}
