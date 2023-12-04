package utils

import "testing"

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
	setInt := NewSet[int]()
	intToRemove := []int{1, 2, 3}
	setInt.Add(intToRemove...)

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
