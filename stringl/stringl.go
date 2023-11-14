package stringl

import (
	"stringlist/list"
)

type StrOnList struct {
	list *list.List
}

// New создает новую строку на списке
func New(s string) *StrOnList {
	l := list.NewList()
	for _, v := range s {
		l.Add(v)
	}
	return &StrOnList{
		list: l,
	}
}

// Len возвращает длину строки на списке
func (sol *StrOnList) Len() int64 {
	return sol.list.Len()
}

// Inject вставляет другую строку на списке в эту строку на списке по индексу
func (sol *StrOnList) Inject(other *StrOnList, index int64) (ok bool) {
	if sol.list.Len() == 0 {
		return false
	}
	if other.list.Len() == 0 {
		return false
	}
	return sol.list.Inject(other.list, index)
}

// Append добавляет rune в конец строки
func (sol *StrOnList) Append(r rune) {
	sol.list.Add(r)
}

// Prepend добавляет rune в начало строки
func (sol *StrOnList) Prepend(r rune) {
	n := list.NewNode(r, nil)
	sol.list.Insert(n, 0)
}

// Insert вставляет rune по индексу
func (sol *StrOnList) Insert(r rune, index int64) (ok bool) {
	n := list.NewNode(r, nil)
	return sol.list.Insert(n, index)
}

// Remove удаляет rune по индексу
func (sol *StrOnList) Remove(index int64) (ok bool) {
	if sol.list.Len() == 0 {
		return false
	}
	return sol.list.RemoveByIndex(index)
}

// String возвращает строковое представление строки на списке
func (sol *StrOnList) String() string {
	s, ok := sol.list.GetAll()
	if !ok {
		return ""
	}
	return string(s)
}

// At возвращает rune по индексу
func (sol *StrOnList) At(index int64) (value rune, ok bool) {
	return sol.list.GetValueByIndex(index)
}

// Concat конкатенирует две строки
func (sol *StrOnList) Concat(other *StrOnList) *StrOnList {
	if other.Len() == 0 {
		return sol
	}
	if sol.Len() == 0 {
		return other
	}
	s1 := New(sol.String())
	ok := s1.Inject(other, s1.Len())
	if !ok {
		return nil
	}
	return s1
}

// Equals проверяет равенство двух строк
func (sol *StrOnList) Equals(other *StrOnList) bool {
	return sol.list.EqualValues(other.list)
}

// IndexOf возвращает индекс первого вхождения rune в строку
func (sol *StrOnList) IndexOf(r rune) int64 {
	return sol.list.GetIndexByValue(r)
}

// Substring возвращает подстроку по индексам от и до (не включая)
func (sol *StrOnList) Substring(from, to int64) *StrOnList {
	sublist := sol.list.GetSubList(from, to)
	if sublist == nil {
		return nil
	}
	return &StrOnList{
		list: sublist,
	}
}

// ReplaceAll заменяет все rune в строке
func (sol *StrOnList) ReplaceAll(old, new rune) {
	for i := int64(0); i < sol.Len(); i++ {
		v, ok := sol.At(i)
		if ok && v == old {
			sol.Remove(i)
			sol.Insert(new, i)
		}
	}
}

// ReplaceOnce заменяет первое вхождение rune в строке
func (sol *StrOnList) ReplaceOnce(old, new rune) {
	for i := int64(0); i < sol.Len(); i++ {
		v, ok := sol.At(i)
		if ok && v == old {
			sol.Remove(i)
			sol.Insert(new, i)
			return
		}
	}
}
