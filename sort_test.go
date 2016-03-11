package sort_test

import (
	"math/rand"
	"testing"
	"testing/quick"

	original "sort"

	"github.com/pcasaretto/sort"
)

func testSortFunc(f sort.SortFunc, t *testing.T) {
	assertion := func(slice []int) bool {
		f(slice)
		for i := 1; i < len(slice); i++ {
			if slice[i] < slice[i-1] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestInsertSort(t *testing.T) { testSortFunc(sort.InsertSort, t) }
func TestBubbleSort(t *testing.T) { testSortFunc(sort.BubbleSort, t) }
func TestMergeSort(t *testing.T)  { testSortFunc(sort.MergeSort, t) }
func TestQuickSort(t *testing.T)  { testSortFunc(sort.QuickSort, t) }

func benchmarkSortFunc(f sort.SortFunc, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(rand.Perm(size))
	}
}

func BenchmarkReference10(b *testing.B)   { benchmarkSortFunc(original.Ints, 10, b) }
func BenchmarkInsertSort10(b *testing.B)  { benchmarkSortFunc(sort.InsertSort, 10, b) }
func BenchmarkBubbleSort10(b *testing.B)  { benchmarkSortFunc(sort.BubbleSort, 10, b) }
func BenchmarkMergeSort10(b *testing.B)   { benchmarkSortFunc(sort.MergeSort, 10, b) }
func BenchmarkQuickSort10(b *testing.B)   { benchmarkSortFunc(sort.QuickSort, 10, b) }
func BenchmarkReference100(b *testing.B)  { benchmarkSortFunc(original.Ints, 100, b) }
func BenchmarkInsertSort100(b *testing.B) { benchmarkSortFunc(sort.InsertSort, 100, b) }
func BenchmarkBubbleSort100(b *testing.B) { benchmarkSortFunc(sort.BubbleSort, 100, b) }
func BenchmarkMergeSort100(b *testing.B)  { benchmarkSortFunc(sort.MergeSort, 100, b) }
func BenchmarkQuickSort100(b *testing.B)  { benchmarkSortFunc(sort.QuickSort, 100, b) }
