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
