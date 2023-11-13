package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode('a', nil)
	assert.Equal(t, 'a', n.value)
	assert.Nil(t, n.nextNode)
}

func TestNewNode2(t *testing.T) {
	n := NewNode('a', &Node{value: 'b'})
	assert.Equal(t, 'a', n.value)
	assert.Equal(t, 'b', n.nextNode.value)
}

func TestNewList(t *testing.T) {
	l := NewList()
	assert.Equal(t, int64(0), l.Len())
	assert.Nil(t, l.firstNode)
	assert.Nil(t, l.lastNode)
}

func TestList_Add(t *testing.T) {
	l := NewList()
	assert.Equal(t, int64(0), l.Len())
	assert.Nil(t, l.firstNode)
	assert.Nil(t, l.lastNode)
	l.Add('a')
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, 'a', l.firstNode.value)
	assert.Equal(t, 'a', l.lastNode.value)
	l.Add('b')
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, 'a', l.firstNode.value)
	assert.Equal(t, 'b', l.lastNode.value)
}

func TestList_RemoveByIndex(t *testing.T) {
	l := NewList()
	l.Add('a')
	l.Add('b')
	l.Add('c')
	l.RemoveByIndex(1)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, 'a', l.firstNode.value)
	assert.Equal(t, 'c', l.lastNode.value)
	l.RemoveByIndex(0)
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, 'c', l.firstNode.value)
}

func TestList_GetValueByIndex(t *testing.T) {
	l := NewList()
	l.Add('a')
	l.Add('b')
	l.Add('c')
	l.Add('d')
	r, ok := l.GetValueByIndex(2)
	assert.True(t, ok)
	assert.Equal(t, 'c', r)
	r, ok = l.GetValueByIndex(4)
	assert.False(t, ok)
	assert.Equal(t, rune(0), r)
}

func TestList_Insert(t *testing.T) {
	l := NewList()
	l.Add('a')
	l.Add('b')
	l.Add('c')
	l.Add('d')
	l.Insert(&Node{value: 'e'}, 1)
	assert.Equal(t, int64(5), l.Len())
	assert.Equal(t, 'a', l.firstNode.value)
	assert.Equal(t, 'e', l.firstNode.nextNode.value)
	assert.Equal(t, 'b', l.firstNode.nextNode.nextNode.value)
	assert.Equal(t, 'c', l.firstNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'd', l.firstNode.nextNode.nextNode.nextNode.nextNode.value)
}

func TestList_Inject(t *testing.T) {
	l := NewList()
	l.Add('a')
	l.Add('b')
	l.Add('c')
	l.Add('d')
	ll := NewList()
	ll.Add('e')
	ok := l.Inject(ll, 1)
	assert.True(t, ok)
	assert.Equal(t, int64(5), l.Len())
	assert.Equal(t, 'a', l.firstNode.value)
	assert.Equal(t, 'e', l.firstNode.nextNode.value)
	assert.Equal(t, 'b', l.firstNode.nextNode.nextNode.value)
	assert.Equal(t, 'c', l.firstNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'd', l.firstNode.nextNode.nextNode.nextNode.nextNode.value)
}

func TestList_Inject2(t *testing.T) {
	l1 := NewList()
	l1.Add('a')
	l1.Add('b')
	l1.Add('c')
	l2 := NewList()
	l2.Add('d')
	l2.Add('e')
	l2.Add('f')
	l1.Inject(l2, 1)
	assert.Equal(t, int64(6), l1.Len())
	assert.Equal(t, 'a', l1.firstNode.value)
	assert.Equal(t, 'd', l1.firstNode.nextNode.value)
	assert.Equal(t, 'e', l1.firstNode.nextNode.nextNode.value)
	assert.Equal(t, 'f', l1.firstNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'b', l1.firstNode.nextNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'c', l1.firstNode.nextNode.nextNode.nextNode.nextNode.nextNode.value)
}

func TestList_Inject3(t *testing.T) {
	l1 := NewList()
	l1.Add('a')
	l1.Add('b')
	l1.Add('c')
	l2 := NewList()
	l2.Add('d')
	l2.Add('e')
	l2.Add('f')
	l1.Inject(l2, 1)
	l1.Inject(l2, 5)
	assert.Equal(t, int64(9), l1.Len())
	assert.Equal(t, 'a', l1.firstNode.value)
	assert.Equal(t, 'd', l1.firstNode.nextNode.value)
	assert.Equal(t, 'e', l1.firstNode.nextNode.nextNode.value)
	assert.Equal(t, 'f', l1.firstNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'b', l1.firstNode.nextNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'd', l1.firstNode.nextNode.nextNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'e', l1.firstNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'f', l1.firstNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.value)
	assert.Equal(t, 'c', l1.firstNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.nextNode.value)
}
