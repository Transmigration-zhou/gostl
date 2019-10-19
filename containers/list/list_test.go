package list

import "testing"

func TestList(t *testing.T) {
	list := New()
	list.PushBack(1)
	if list.Front().Value != 1 || list.Back().Value != 1 {
		t.Errorf("list data error: %d %d", list.Front().Value, list.Back().Value)
	}
	list.PushFront(2)

	if list.Len() != 2 {
		t.Errorf("list len error: %d", list.Len())
	}
	list.PushBack(3)
	list.PushFront(4)
	if list.Front().Value != 4 || list.Back().Value != 3 {
		t.Errorf("list data error: %d %d", list.Front().Value, list.Back().Value)
	}
	t.Logf("list: %v", list)

	list.Remove(list.Front())
	t.Logf("list: %v", list)
	if list.String() != "[2 1 3]" {
		t.Errorf("list data error: %s", list.String())
	}

	list.Remove(list.Back())
	if list.String() != "[2 1]" {
		t.Errorf("list data error: %s", list.String())
	}
	list.PushBack(5)
	list.PushBack(6)
	list.InsertAfter(7, list.Front())
	t.Logf("list: %v", list)
	if list.String() != "[2 7 1 5 6]" {
		t.Errorf("list data error: %s", list.String())
	}
	list.InsertBefore(8, list.Back().Prev())
	t.Logf("list: %v", list)
	if list.String() != "[2 7 1 8 5 6]" {
		t.Errorf("list data error: %s", list.String())
	}

	list.Remove(list.Front().Next().Next())
	t.Logf("list: %v", list)
	if list.String() != "[2 7 8 5 6]" {
		t.Errorf("list data error: %s", list.String())
	}

	list.Remove(list.Front())
	t.Logf("list: %v", list)
	if list.String() != "[7 8 5 6]" {
		t.Errorf("list data error: %s", list.String())
	}

	list.Remove(list.Back())
	t.Logf("list: %v", list)
	if list.String() != "[7 8 5]" {
		t.Errorf("list data error: %s", list.String())
	}

	list.PushBackList(list)
	t.Logf("list: %v", list)
	if list.String() != "[7 8 5 7 8 5]" {
		t.Errorf("list data error: %s", list.String())
	}
	/////////////////////////////
	list2 := New()
	list2.PushBack(1)
	list2.PushBack(2)
	list2.PushBack(3)
	list2.PushFrontList(list2)
	t.Logf("list: %v", list2)
	if list2.String() != "[1 2 3 1 2 3]" {
		t.Errorf("list data error: %s", list2.String())
	}

	list.PushBackList(list2)
	t.Logf("list: %v", list)
	if list.String() != "[7 8 5 7 8 5 1 2 3 1 2 3]" || list.Len() != 12 {
		t.Errorf("list data error: %s", list.String())
	}
}
