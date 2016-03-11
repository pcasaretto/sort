package sort

type SortFunc func([]int)

func InsertSort(in []int) {
	var i int
	for j := 1; j < len(in); j++ {
		key := in[j]
		i = j - 1
		for i > -1 && in[i] > key {
			in[i+1] = in[i]
			i = i - 1
		}
		in[i+1] = key
	}
}

func BubbleSort(in []int) {
	sorted := true
	var n0, n1 int
	for sorted {
		sorted = false
		for i := 1; i < len(in); i++ {
			n0 = in[i-1]
			n1 = in[i]
			if n0 > n1 {
				in[i] = n0
				in[i-1] = n1
				sorted = true
			}
		}
	}
}

func QuickSort(in []int) {
	if len(in) <= 1 {
		return
	}
	p := in[len(in)-1]
	var i int
	for j, v := range in {
		if v < p {
			in[j] = in[i]
			in[i] = v
			i++
		}
	}
	v := in[i]
	in[i] = p
	in[len(in)-1] = v

	QuickSort(in[:i])
	QuickSort(in[i:])

	return
}

func MergeSort(in []int) {
	q := len(in) / 2
	if q == 0 {
		return
	}
	left := in[0:q]
	right := in[q:len(in)]
	MergeSort(left)
	MergeSort(right)

	// merge left and right
	leftTemp := make([]int, len(left))
	copy(leftTemp, left)
	rightTemp := make([]int, len(right))
	copy(rightTemp, right)

	var i, j, k int
	for i < len(leftTemp) && j < len(rightTemp) {
		if leftTemp[i] < rightTemp[j] {
			in[k] = leftTemp[i]
			i++
		} else {
			in[k] = rightTemp[j]
			j++
		}
		k++
	}

	for i < len(leftTemp) {
		in[k] = leftTemp[i]
		i++
		k++
	}
	for j < len(rightTemp) {
		in[k] = rightTemp[j]
		j++
		k++
	}

}
