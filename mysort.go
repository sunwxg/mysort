package mysort

import "reflect"

//Compare to increasing order: a < b return true
//
//Compare to decreasing order: a > b return true
type Compare func(a, b interface{}) bool

func swapValue(a, b reflect.Value) {
	t := a.Interface()
	a.Set(b)
	b.Set(reflect.ValueOf(t))
}

func InsertSort(l interface{}, f Compare) {

	list := reflect.ValueOf(l)
	if list.Kind() != reflect.Slice {
		return
	}

	for i := 1; i < list.Len(); i++ {
		for k := i; k > 0; k-- {
			a := list.Index(k - 1).Interface()
			b := list.Index(k).Interface()
			if f(a, b) {
				break
			} else {
				swapValue(list.Index(k-1), list.Index(k))
			}
		}
	}
}

func selectIndex(list reflect.Value, left, right int) int {
	return left
}

func partition(list reflect.Value, left, right int, f Compare) int {

	index := selectIndex(list, left, right)

	storeIndex := left

	swapValue(list.Index(index), list.Index(right))

	b := list.Index(right).Interface()
	for i := left; i < right; i++ {
		a := list.Index(i).Interface()
		if f(a, b) {
			swapValue(list.Index(i), list.Index(storeIndex))
			storeIndex++
		}
	}

	swapValue(list.Index(right), list.Index(storeIndex))
	return storeIndex
}

func qSort(list reflect.Value, left, right int, f Compare) {

	if left >= right {
		return
	}

	index := partition(list, left, right, f)

	qSort(list, left, index-1, f)
	qSort(list, index+1, right, f)
}

func QuickSort(l interface{}, f Compare) {
	list := reflect.ValueOf(l)
	if list.Kind() != reflect.Slice {
		return
	}

	qSort(list, 0, list.Len()-1, f)
}

func shiftUp(list reflect.Value, i int, f Compare) {

	if i == 0 {
		return
	}

	iParent := (i - 1) / 2

	a := list.Index(iParent).Interface()
	b := list.Index(i).Interface()
	if f(a, b) {
		swapValue(list.Index(iParent), list.Index(i))
		shiftUp(list, iParent, f)
	}
}

func buildHeap(list reflect.Value, f Compare) {
	for i := 1; i < list.Len(); i++ {
		shiftUp(list, i, f)
	}
}

func shiftDown(list reflect.Value, i, iMax int, f Compare) {

	iLeft := 2*i + 1
	if iLeft > iMax {
		return
	} else {
		a := list.Index(i).Interface()
		b := list.Index(iLeft).Interface()
		if f(a, b) {
			swapValue(list.Index(iLeft), list.Index(i))
			shiftDown(list, iLeft, iMax, f)
		}
	}

	iRight := 2*i + 2
	if iRight > iMax {
		return
	} else {
		a := list.Index(i).Interface()
		b := list.Index(iRight).Interface()
		if f(a, b) {
			swapValue(list.Index(iRight), list.Index(i))
			shiftDown(list, iRight, iMax, f)
		}
	}
}

func HeapSort(l interface{}, f Compare) {
	list := reflect.ValueOf(l)
	if list.Kind() != reflect.Slice {
		return
	}

	buildHeap(list, f)

	for i := list.Len() - 1; i > 0; i-- {
		swapValue(list.Index(0), list.Index(i))
		shiftDown(list, 0, i-1, f)
	}
}
