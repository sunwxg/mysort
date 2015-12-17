package mysort_test

import (
	"math/rand"
	"testing"

	"github.com/sunwxg/mysort"
)

//a > b = ture
func c(a, b interface{}) bool {
	av := a.(int)
	bv := b.(int)

	switch {
	case av > bv:
		return true
	case av <= bv:
		return false
	}
	return false
}

//a < b = ture
func u(a, b interface{}) bool {
	av := a.(int)
	bv := b.(int)

	switch {
	case av < bv:
		return true
	case av >= bv:
		return false
	}
	return false
}

func TestInsertSort(t *testing.T) {
	list := []int{1, 2, 3, 6, 5}
	mysort.InsertSort(list, c)

	result := list
	wanted := []int{6, 5, 3, 2, 1}

	for i, v := range wanted {
		if v != result[i] {
			t.Fatalf("InsertSort: wanted (%v), geted (%v)", wanted, result)
			break
		}
	}
}

func TestQuickSort(t *testing.T) {
	list := []int{1, 2, 3, 6, 5}
	mysort.QuickSort(list, c)

	result := list
	wanted := []int{6, 5, 3, 2, 1}

	for i, v := range wanted {
		if v != result[i] {
			t.Fatalf("QuickSort: wanted (%v), geted (%v)", wanted, result)
			break
		}
	}
}

func TestHeapSortDown(t *testing.T) {
	list := []int{1, 2, 3, 6, 5}
	mysort.HeapSort(list, c)

	result := list
	wanted := []int{6, 5, 3, 2, 1}

	for i, v := range wanted {
		if v != result[i] {
			t.Fatalf("HeapSort: wanted (%v), geted (%v)", wanted, result)
			break
		}
	}
}

func TestHeapSortUp(t *testing.T) {
	list := []int{9, 2, 3, 6, 5}
	mysort.HeapSort(list, u)

	result := list
	wanted := []int{2, 3, 5, 6, 9}

	for i, v := range wanted {
		if v != result[i] {
			t.Fatalf("HeapSort: wanted (%v), geted (%v)", wanted, result)
			break
		}
	}
}

func TestDown(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	var l []int
	for i := 0; i < 100; i++ {
		l = append(l, r.Intn(1000))
	}

	listInsert := make([]int, 100)
	for i, _ := range listInsert {
		listInsert[i] = l[i]
	}
	mysort.InsertSort(listInsert, c)

	listQuickSort := make([]int, 100)
	for i, _ := range listQuickSort {
		listQuickSort[i] = l[i]
	}
	mysort.QuickSort(listQuickSort, c)

	listHeapSort := make([]int, 100)
	for i, _ := range listHeapSort {
		listHeapSort[i] = l[i]
	}
	mysort.HeapSort(listHeapSort, c)

	for i, v := range listInsert {
		if v != listQuickSort[i] {
			t.Fatalf("HeapSort don't equal QuickSort")
			break
		}
	}

	for i, v := range listInsert {
		if v != listHeapSort[i] {
			t.Fatalf("HeapSort don't equal HeapSort")
			break
		}
	}
}

func TestUp(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	var l []int
	for i := 0; i < 100; i++ {
		l = append(l, r.Intn(1000))
	}

	listInsert := make([]int, 100)
	for i, _ := range listInsert {
		listInsert[i] = l[i]
	}
	mysort.InsertSort(listInsert, u)

	listQuickSort := make([]int, 100)
	for i, _ := range listQuickSort {
		listQuickSort[i] = l[i]
	}
	mysort.QuickSort(listQuickSort, u)

	listHeapSort := make([]int, 100)
	for i, _ := range listHeapSort {
		listHeapSort[i] = l[i]
	}
	mysort.HeapSort(listHeapSort, u)

	for i, v := range listInsert {
		if v != listQuickSort[i] {
			t.Fatalf("HeapSort don't equal QuickSort")
			break
		}
	}

	for i, v := range listInsert {
		if v != listHeapSort[i] {
			t.Fatalf("HeapSort don't equal HeapSort")
			break
		}
	}
}
