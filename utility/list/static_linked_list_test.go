package list

import (
	"testing"
)

func TestFitnessStaticLinkedList(t *testing.T) {
	lList := NewStaticLinkedList(3)

	fitness1, fitness2, fitness3 := 11, 22, 33

	// 0
	err := lList.Add(fitness1, fitness2)
	if err != nil {
		t.Error(err)
	}

	// 1
	err = lList.Add(fitness2, fitness3)
	if err != nil {
		t.Error(err)
	}

	// 2
	err = lList.Add(fitness1, fitness3)
	if err != nil {
		t.Error(err)
	}

	// err
	err = lList.Add(fitness3, fitness1)
	if err == nil {
		t.Errorf("%v", lList)
	}

	// 2 true
	idx1, flag := lList.GetFirstDataIndex(fitness1)
	if idx1 != 2 || !flag {
		t.Errorf("fitness1 %v %v", idx1, flag)
	}

	// 1 true
	idx2, flag := lList.GetFirstDataIndex(fitness2)
	if idx2 != 1 || !flag {
		t.Errorf("fitness2 %v %v", idx2, flag)
	}

	// 0 false
	idx3, flag := lList.GetFirstDataIndex(fitness3)
	if idx3 != 0 || flag {
		t.Errorf("fitness1 %v %v", idx3, flag)
	}

	// fitness3 0
	fit1, next := lList.GetData(idx1)
	if next != 0 {
		t.Errorf("fitness1 %v %v", fit1, next)
	}

	// fitness2 -1
	fit1, next = lList.GetData(next)
	if next != -1 {
		t.Errorf("fitness1 %v %v", fit1, next)
	}

	// fitness3 -1
	fit2, next := lList.GetData(idx2)
	if next != -1 {
		t.Errorf("fitness2 %v %v", fit2, next)
	}

	// 0 -1
	fit3, next := lList.GetData(-1)
	if fit3 != 0 || next != -1 {
		t.Errorf("fitness3 %v %v", fit3, next)
	}

	if size1, exist := lList.GetSize(fitness1); size1 != 2 || !exist {
		t.Errorf("fitness1 %v %v", size1, exist)
	}
}
