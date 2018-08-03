package list

import (
	"fmt"
)

// StaticLinkedList is a linked list which saves the relationship of a block of memory
type StaticLinkedList struct {
	// head saves head of linked list which belong to different Fitness
	head map[interface{}]int
	// size saves the length of different linked list
	size map[interface{}]int
	// datas saves data
	datas []int
	// next saves index of the next element of current element
	next []int
	// total saves the amount of elements are used
	total int
	// cap is the capacity
	cap int
}

// NewStaticLinkedList returns a *fitnessStaticLinkedList
func NewStaticLinkedList(capacity int) *StaticLinkedList {
	list := &StaticLinkedList{
		head:  make(map[interface{}]int),
		size:  make(map[interface{}]int),
		datas: make([]int, capacity),
		next:  make([]int, capacity),
		total: 0,
		cap:   capacity,
	}
	list.Reset()
	return list
}

// Reset resets the fitnessStaticLinkedList
func (list *StaticLinkedList) Reset() {
	list.head = make(map[interface{}]int)
	list.size = make(map[interface{}]int)
	list.total = 0
}

// Add gets an element isn't used and returns it's index
func (list *StaticLinkedList) Add(headID interface{}, elemIndex int) error {
	if list.cap <= list.total {
		return fmt.Errorf("the capacity of linked list is full")
	}
	loc, flag := list.head[headID]
	newIndex := list.total
	list.total++
	if flag {
		list.next[newIndex] = loc
	} else {
		list.next[newIndex] = -1
	}
	list.datas[newIndex] = elemIndex
	list.head[headID] = newIndex
	list.size[headID]++
	return nil
}

// GetFirstDataIndex returns the first data's index if the headFit exists
func (list *StaticLinkedList) GetFirstDataIndex(headID interface{}) (int, bool) {
	idx, exist := list.head[headID]
	return idx, exist
}

// GetData returns the data, next index of input index, if the index doesn't exist, returns (0,-1)
func (list *StaticLinkedList) GetData(index int) (int, int) {
	if index < 0 || index >= list.total || index >= list.cap {
		return 0, -1
	}
	return list.datas[index], list.next[index]
}

// GetSize returns the length of linked list which head is headID if exists
func (list *StaticLinkedList) GetSize(headID interface{}) (int, bool) {
	amount, exist := list.size[headID]
	return amount, exist
}

// GetHead returns a copy of head
func (list *StaticLinkedList) GetHead() map[interface{}]int {
	ans := make(map[interface{}]int, len(list.head))
	for k, v := range list.head {
		ans[k] = v
	}
	return ans
}
