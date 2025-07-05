package arraylist

import (
	"container/list"
	"fmt"
	"slices"
	"strings"

	"github.com/kianooshaz/ds/types"
)

type List[T comparable] struct {
	elements []T
}

const (
	growthFactor = float64(2.0)
	shrinkFactor = float64(.25)
)

func (list *List[T]) Add(values ...T) {
	l := len(list.elements)
	list.growBy(len(values))

	for i := range values {
		list.elements[l+i] = values[i]
	}
}

func (list *List[T]) Get(index int) (T, bool) {

	if !list.withinRange(index) {
		var t T
		return t, false
	}

	return list.elements[index], true
}

func (list *List[T]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements = slices.Delete(list.elements, index, index+1)
	list.shrink()
}

func (list *List[T]) Contains(values ...T) bool {

	for _, v := range values {
		if !slices.Contains(list.elements, v) {
			return false
		}
	}

	return true
}

func (list *List[T]) Values() []T {
	return slices.Clone(list.elements)
}

func (list *List[T]) IndexOf(value T) int {
	return slices.Index(list.elements, value)
}

func (list *List[T]) Empty() bool {
	return len(list.elements) == 0
}

func (list *List[T]) Size() int {
	return len(list.elements)
}

func (list *List[T]) Clear() {
	clear(list.elements[:cap(list.elements)])
	list.elements = list.elements[:0]
}

func (list *List[T]) Sort(cmp types.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}

	slices.SortFunc(list.elements, cmp)
}

func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		if index == len(list.elements) {
			list.Add(values...)
		}
		return
	}

	l := len(list.elements)
	list.growBy(len(values))
	list.elements = slices.Insert(list.elements[:l], index, values...)
}

func (list *List[T]) Set(index int, value T) {

	if !list.withinRange(index) {
		if index == len(list.elements) {
			list.Add(value)
		}

		return
	}

	list.elements[index] = value
}

func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := make([]string, 0,len(list.elements))
	for _, value := range list.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	
	return str
}

func (list *List[T]) growBy(n int) {
	currentCapacity := cap(list.elements)

	if newLength := len(list.elements) + n; newLength >= currentCapacity {
		newCapacity := int(growthFactor * float64(newLength))
		list.resize(newLength, newCapacity)
	} else {
		list.elements = list.elements[:newLength]
	}
}

func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < len(list.elements)
}

func (list *List[T]) resize(length, capacity int) {
	newElements := make([]T, length, capacity)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List[T]) shrink() {

	if shrinkFactor == 0.0 {
		return
	}

	currentCapacity := cap(list.elements)
	if len(list.elements) <= int(float64(currentCapacity)*shrinkFactor) {
		list.resize(len(list.elements), len(list.elements))
	}
}
