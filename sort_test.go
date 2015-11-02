package sort_test

import (
	"math/rand"
	"reflect"
	"testing"

	original "sort"

	"github.com/pcasaretto/sort"
)

func testSortFunc(f sort.SortFunc, t *testing.T) {
	input := rand.Perm(10)
	expected := make([]int, len(input))
	copy(expected, input)
	original.Ints(expected)
	f(input)
	if !reflect.DeepEqual(expected, input) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

// func testSortFunc(f sort.SortFunc, t *testing.T) {
// 	if err := quick.CheckEqual(f, sort.SortFunc(original.Ints), nil); err != nil {
// 		t.Error(err)
// 	}
// }

func TestInsertSort(t *testing.T) { testSortFunc(sort.InsertSort, t) }
func TestBubbleSort(t *testing.T) { testSortFunc(sort.BubbleSort, t) }

func benchmarkSortFunc(f sort.SortFunc, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(rand.Perm(size))
	}
}

func BenchmarkReference10(b *testing.B)   { benchmarkSortFunc(original.Ints, 10, b) }
func BenchmarkInsertSort10(b *testing.B)  { benchmarkSortFunc(sort.InsertSort, 10, b) }
func BenchmarkBubbleSort10(b *testing.B)  { benchmarkSortFunc(sort.BubbleSort, 10, b) }
func BenchmarkReference100(b *testing.B)  { benchmarkSortFunc(original.Ints, 100, b) }
func BenchmarkInsertSort100(b *testing.B) { benchmarkSortFunc(sort.InsertSort, 100, b) }
func BenchmarkBubbleSort100(b *testing.B) { benchmarkSortFunc(sort.BubbleSort, 100, b) }
