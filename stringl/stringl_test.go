package stringl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	sl := New("Hello, World!")
	assert.Equal(t, "Hello, World!", sl.String())
}

func TestStrOnList_Len(t *testing.T) {
	sl := New("Hello, World!")
	assert.Equal(t, int64(13), sl.Len())
}

func TestStrOnList_Append(t *testing.T) {
	sl := New("Hello, World")
	sl.Append('!')
	assert.Equal(t, "Hello, World!", sl.String())
	assert.Equal(t, int64(13), sl.Len())
}

func TestStrOnList_Prepend(t *testing.T) {
	sl := New("Hello, World")
	sl.Prepend('!')
	assert.Equal(t, "!Hello, World", sl.String())
	assert.Equal(t, int64(13), sl.Len())
}

func TestStrOnList_Insert(t *testing.T) {
	sl := New("Hello, World")
	sl.Insert('!', 6)
	assert.Equal(t, "Hello,! World", sl.String())
	assert.Equal(t, int64(13), sl.Len())
}

func TestStrOnList_Remove(t *testing.T) {
	sl := New("Hello, World")
	sl.Remove(5)
	assert.Equal(t, "Hello World", sl.String())
	assert.Equal(t, int64(11), sl.Len())
}

func TestStrOnList_String(t *testing.T) {
	sl := New("Hello, World!")
	assert.Equal(t, "Hello, World!", sl.String())
}

func TestStrOnList_At(t *testing.T) {
	sl := New("Hello")
	r, ok := sl.At(0)
	assert.Equal(t, 'H', r)
	assert.True(t, ok)
	r, ok = sl.At(1)
	assert.Equal(t, 'e', r)
	assert.True(t, ok)
	r, ok = sl.At(2)
	assert.Equal(t, 'l', r)
	assert.True(t, ok)
	r, ok = sl.At(3)
	assert.Equal(t, 'l', r)
	assert.True(t, ok)
	r, ok = sl.At(4)
	assert.Equal(t, 'o', r)
	assert.True(t, ok)
	r, ok = sl.At(5)
	assert.Equal(t, rune(0), r)
	assert.False(t, ok)
}

func TestStrOnList_Concat(t *testing.T) {
	sl1 := New("Hello, ")
	sl2 := New("World!")
	sl3 := sl1.Concat(sl2)
	assert.Equal(t, "Hello, World!", sl3.String())
	assert.Equal(t, int64(13), sl3.Len())
}

func TestStrOnList_Equals(t *testing.T) {
	sl1 := New("Hello, World!")
	sl2 := New("Hello, World!")
	assert.True(t, sl1.Equals(sl2))
	sl3 := New("Hello, World")
	assert.False(t, sl1.Equals(sl3))
}

func TestStrOnList_IndexOf(t *testing.T) {
	sl := New("Hello")
	assert.Equal(t, int64(0), sl.IndexOf('H'))
	assert.Equal(t, int64(1), sl.IndexOf('e'))
	assert.Equal(t, int64(2), sl.IndexOf('l'))
	assert.Equal(t, int64(2), sl.IndexOf('l'))
	assert.Equal(t, int64(4), sl.IndexOf('o'))
	assert.Equal(t, int64(-1), sl.IndexOf('!'))
}

func TestStrOnList_Substring(t *testing.T) {
	sl := New("Hello, World!")
	assert.Equal(t, "Hello", sl.Substring(0, 5).String())
	assert.Equal(t, "World", sl.Substring(7, 12).String())
}

func TestStrOnList_ReplaceAll(t *testing.T) {
	sl := New("Hello, World!")
	sl.ReplaceAll('o', '0')
	assert.Equal(t, "Hell0, W0rld!", sl.String())
	assert.Equal(t, int64(13), sl.Len())
}

func TestStrOnList_ReplaceOnce(t *testing.T) {
	sl := New("Hello, World!")
	sl.ReplaceOnce('l', '1')
	assert.Equal(t, "He1lo, World!", sl.String())
	assert.Equal(t, int64(13), sl.Len())
}
