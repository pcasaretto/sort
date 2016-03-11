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
	copy(in, quickSort(in))
}

func quickSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}
	p := in[0]
	left := make([]int, 0, len(in))
	right := make([]int, 0, len(in))
	eq := make([]int, 0, len(in))
	for _, v := range in {
		switch {
		case v < p:
			left = append(left, v)
		case v == p:
			eq = append(eq, v)
		case v > p:
			right = append(right, v)
		}
	}

	return append(quickSort(left), append(quickSort(eq), quickSort(right)...)...)
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
