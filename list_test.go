package list

import "testing"

func TestNewList(t *testing.T) {
	l := New()
	if l == nil {
		t.Fatal("should return new list")
	}
	if l.Size() != 0 {
		t.Fatal("should return length zero")
	}
	if l.First() != nil {
		t.Fatal("should return nil head")
	}
	if l.Last() != nil {
		t.Fatal("should return nil tail")
	}
}

func TestListAdd(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	if l.Last().value != 99 {
		t.Fatal("last element should be 99")
	}
}

func TestListPop(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	for i := 0; i < 100; i++ {
		v, _ := l.Pop()
		if v != i {
			t.Fatal("should return the same value")
		}
	}
	if l.Size() != 0 {
		t.Fatal("should be length zero")
	}
	if l.First() != nil {
		t.Fatal("should be nil")
	}
	if l.Last() != nil {
		t.Fatal("should be nil")
	}
	v, err := l.Pop()
	if err == nil {
		t.Fatal("should not be nil")
	}
	if v != 0 {
		t.Fatal("should be zero")
	}
}

func TestListPush(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Push(i)
		if l.First().value != i {
			t.Fatal("should be the same value")
		}
	}
	if l.Size() != 100 {
		t.Fatal("should have a length ninty-nine")
	}
	if l.Last().value != 0 {
		t.Fatal("element should contain value zero")
	}
	if l.First().value != 99 {
		t.Fatal("should contain value of ninty-nine")
	}
}

func TestListFind(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	v, err := l.Find(20)
	if err != nil {
		t.Fatal("should find this value")
	}
	if v.Value() != 20 {
		t.Fatal("should be the same value")
	}
	v, err = l.Find(99999)
	if err == nil {
		t.Fatal("should not find this value")
	}
	if v != nil {
		t.Fatal("should not exist")
	}
}

func TestListValues(t *testing.T) {
	l := New()
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	av := l.Values()
	if av == nil {
		t.Fatal("should not be nil")
	}
	for i := 0; i < len(want); i++ {
		if av[i] != want[i] {
			t.Fatal("should be the same value")
		}
	}
}

func BenchmarkListFind(b *testing.B) {
	l := New()
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	l.Find(2000)
}
