package list

import (
	"math"
	"testing"
)

func TestNewList(t *testing.T) {
	l := New()
	if l == nil {
		t.Fatal("should return new list")
	}
	if l.Length() != 0 {
		t.Fatal("should return length zero")
	}
	if l.Head() != nil {
		t.Fatal("should return nil head")
	}
	if l.Head() != nil {
		t.Fatal("should return nil tail")
	}
}
func TestListAdd(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	if l.Tail().value != 99 {
		t.Fatal("last element should be 99")
	}
}
func TestListPop(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	for i := 0; i < 100; i++ {
		v := l.Pop()
		if v.Value() != i {
			t.Fatal("should return the same value")
		}
	}
	if l.Length() != 0 {
		t.Fatal("should be length zero")
	}
	if l.Head() != nil {
		t.Fatal("should be nil")
	}
	if l.Tail() != nil {
		t.Fatal("should be nil")
	}
	v := l.Pop()
	if v != nil {
		t.Fatal("should not be nil")
	}
}
func TestListPush(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Push(i)
		if l.Head().value != i {
			t.Fatal("should be the same value")
		}
	}
	if l.Length() != 100 {
		t.Fatal("should have a length ninty-nine")
	}
	if l.Tail().value != 0 {
		t.Fatal("element should contain value zero")
	}
	if l.Head().value != 99 {
		t.Fatal("should contain value of ninty-nine")
	}
}
func TestListFind(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	v := l.Find(20)
	if v == nil {
		t.Fatal("should find this value")
	}
	if v.Value() != 20 {
		t.Fatal("should be the same value")
	}
	v = l.Find(99999)
	if v != nil {
		t.Fatal("should not find this value")
	}
}
func TestListToSlice(t *testing.T) {
	l := New()
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	av := l.ToSlice()
	if av == nil {
		t.Fatal("should not be nil")
	}
	for i := 0; i < len(want); i++ {
		if av[i] != want[i] {
			t.Fatal("should be the same value")
		}
	}
}
func TestListGetNextNode(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	n := l.Head().Next()
	if n.Value() != 1 {
		t.Fatal("second found should contain a value of one")
	}
}
func TestListTraverseTo(t *testing.T) {
	var n *Node
	l := New()
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	n = l.Traverse(2)
	if n.Value() != 2 {
		t.Fatal("The value at index two should be three")
	}
	n = l.Traverse(math.MinInt32)
	if n != nil {
		t.Fatal("The node should not exist")
	}
	n = l.Traverse(math.MaxInt64)
	if n != nil {
		t.Fatal("The node should not exist")
	}
	n = l.Traverse(9)
	if n.Value() != 9 {
		t.Fatal("The value at this index should be ninty-nine")
	}
	n = l.Traverse(l.Length())
	if n != nil {
		t.Fatal("the node should not exist")
	}
}
func TestListInsert(t *testing.T) {
	l := New()
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	l.Insert(math.MaxInt32, 5)
	if l.Traverse(5).value != math.MaxInt32 {
		t.Fatal("The node should contain the inserted value")
	}
	if l.Insert(math.MaxInt32, math.MinInt32) != nil {
		t.Fatal("The node should not be inserted")
	}
	if l.Insert(math.MaxInt32, math.MaxInt32) != nil {
		t.Fatal("The node should not be inserted")
	}
}
func TestListRemove(t *testing.T) {
	l := New()
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	if l.Remove(math.MinInt32) != nil {
		t.Fatal("should not be removed")
	}
	if l.Remove(math.MaxInt32) != nil {
		t.Fatal("should not be removed")
	}
	if l.Remove(0).value != 0 {
		t.Fatal("should remove the head")
	}
	if l.Remove(l.length-1).value != 9 {
		t.Fatal("should remove the tail node")
	}
	if l.Remove(l.length) != nil {
		t.Fatal("node should not exist")
	}
}
func TestListReverse(t *testing.T) {
	l := New()
	for i := 0; i < 4; i++ {
		l.Add(i)
	}
	l.Reverse()
}
func BenchmarkListFind(b *testing.B) {
	l := New()
	for i := 0; i < 10000; i++ {
		l.Add(i)
	}
	l.Find(2000)
}
