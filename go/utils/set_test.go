package utils

import (
	"reflect"
	"testing"
)

func TestAddSetElements(t *testing.T) {
	intToAdd := []int{1, 2, 3}

	t.Run("int test one by one", func(t *testing.T) {
		setInt := NewSet[int]()
		for _, elm := range intToAdd {
			setInt.Add(elm)
			if !setInt.Exists(elm) {
				t.Errorf("should have been able to add %d to set %+v", elm, setInt)
			}
		}
	})
	t.Run("int test one by one", func(t *testing.T) {
		setInt := NewSet[int]()
		setInt.Add(intToAdd...)
		for _, elm := range intToAdd {
			if !setInt.Exists(elm) {
				t.Errorf("should have been able to add %d to set %+v", elm, setInt)
			}
		}
	})
}

func TestRemoveElements(t *testing.T) {
	setInt := NewSet[int](1, 2, 3)

	if ok := setInt.Remove(1); ok && setInt.Exists(1) {
		t.Errorf("should have removed element %d from %+v", 1, setInt)
	} else if ok != true {
		t.Errorf("should have been able to remove %d from %+v", 1, setInt)
	}
	if !setInt.Exists(2) {
		t.Errorf("should have kept element %d from %+v", 2, setInt)
	}
	if ok := setInt.Remove(1); ok {
		t.Errorf("should not have been able to remove %d again from %+v", 1, setInt)
	}
}

func TestSetLen(t *testing.T) {
	setInt := NewSet[int]()
	intToAdd := []int{1, 2, 3}
	for i, element := range intToAdd {
		if setInt.Len() != i {
			t.Errorf("length %d of set should have been %d", setInt.Len(), i)
		}
		setInt.Add(element)
	}
}

func TestSetUnion(t *testing.T) {
	setA := NewSet[int](1, 2, 3)
	setB := NewSet[int](3, 4, 5)
	setExpected := NewSet[int](1, 2, 3, 4, 5)

	result := setA.Union(setB)
	if !reflect.DeepEqual(result, setExpected) {
		t.Errorf("expected %+v, got %+v", setExpected, result)
	}
}

func TestSetIntersect(t *testing.T) {
	setA := NewSet[int](1, 2, 3)
	setB := NewSet[int](1, 3, 5)
	setC := NewSet[int](2, 3, 7)
	setExpect := NewSet[int](3)
	result := setA.Intersection(setB, setC)
	if !reflect.DeepEqual(result, setExpect) {
		t.Errorf("expected %+v, got %+v", setExpect, result)
	}
	setExpect = NewSet[int](1, 3)
	result = setA.Intersection(setB)
	if !reflect.DeepEqual(result, setExpect) {
		t.Errorf("expected %+v, got %+v", setExpect, result)
	}
}
