package list

type Node struct {
	value    rune
	nextNode *Node
}

func NewNode(value rune, nextNode *Node) (n *Node) {
	return &Node{value: value, nextNode: nextNode}
}

type List struct {
	length    int64 // Текущая длина списка (количество узлов)
	firstNode *Node // Указатель на первый узел
	lastNode  *Node // Указатель на последний узел (для ускорения вставки элемента в конец)
}

// NewList создает новый пустой односвязный список
func NewList() (l *List) {
	return &List{length: 0, firstNode: nil, lastNode: nil}
}

func NewListFromSlice(s []rune) (l *List) {
	l = NewList()
	for _, v := range s {
		l.Add(v)
	}
	return l
}

// Len возвращает количество элементов в списке
func (l *List) Len() int64 {
	return l.length
}

// Add добавляет элемент в конец списка, возвращает индекс добавленного элемента
func (l *List) Add(data rune) (index int64) {
	newNode := &Node{value: data}
	l.length++
	nextNode := l.firstNode
	if nextNode == nil {
		l.firstNode = newNode
		l.lastNode = l.firstNode
		return l.length - 1
	}
	l.lastNode.nextNode = newNode
	l.lastNode = l.lastNode.nextNode
	return l.length - 1
}

// RemoveByIndex удаляет элемент по индексу (текущему порядковому номеру)
func (l *List) RemoveByIndex(index int64) (ok bool) {
	if index >= l.Len() || index < 0 {
		return false
	}
	if index == 0 {
		l.firstNode = l.firstNode.nextNode
		l.length--
		return true
	}
	leftNode := l.GetNodeByIndex(index - 1)
	removableNode := leftNode.nextNode
	if removableNode == nil {
		return false
	}
	leftNode.nextNode = removableNode.nextNode
	l.length--
	return true
}

// GetValueByIndex возвращает значение элемента с данным индексом
func (l *List) GetValueByIndex(index int64) (value rune, ok bool) {
	n := l.GetNodeByIndex(index)
	if n == nil {
		return 0, false
	}
	return n.value, true
}

// GetAll возвращает значения всех элементов
func (l *List) GetAll() (values []rune, ok bool) {
	if l.length == 0 {
		return nil, false
	}

	currentNode := l.firstNode
	currentIndex := int64(0)
	for ; currentNode != l.lastNode.nextNode && currentIndex < l.length; currentNode = currentNode.nextNode {
		values = append(values, currentNode.value)
	}
	return values, true
}

// Clear удаляет все элементы из списка
func (l *List) Clear() {
	l.firstNode = nil
	l.lastNode = nil
	l.length = 0
}

// GetNodeByIndex возвращает указатель на узел с данным индексом
func (l *List) GetNodeByIndex(index int64) (node *Node) {
	if index >= l.Len() || index < 0 {
		return nil
	}
	var currentNode *Node = l.firstNode
	var currentIndex int64 = 0
	for ; currentIndex < index; currentIndex++ {
		currentNode = currentNode.nextNode
	}
	return currentNode
}

// GetIndexByValue возвращает индекс первого найденного элемента с данным значением
func (l *List) GetIndexByValue(value rune) (index int64) {
	currentNode := l.firstNode
	currentIndex := int64(0)
	for ; currentNode != nil; currentNode = currentNode.nextNode {
		if currentNode.value == value {
			return currentIndex
		}
		currentIndex++
	}
	return -1
}

// Insert вставляет Node в список по индексу
func (l *List) Insert(node *Node, index int64) (ok bool) {
	if index > l.Len() || index < 0 {
		return false
	}
	if index == 0 {
		node.nextNode = l.firstNode
		l.firstNode = node
		l.length++
		return true
	}
	if index == l.length {
		l.lastNode.nextNode = node
		l.lastNode = node
		l.length++
		return true
	}
	leftNode := l.GetNodeByIndex(index - 1)
	rightNode := leftNode.nextNode
	leftNode.nextNode = node
	node.nextNode = rightNode
	l.length++
	return true
}

// Inject добавляет другой список в текущий	по индексу
// При угрозе возникновения петель, производится копирование добавляемого списка
func (l *List) Inject(other *List, index int64) (ok bool) {
	if index > l.Len() || index < 0 {
		return false
	}
	var needCopy bool = false
	if l.ContainsNode(other.firstNode) || l.ContainsNode(other.lastNode) {
		needCopy = true
	}
	if needCopy {
		otherSlice, _ := other.GetAll()
		other = NewListFromSlice(otherSlice)
	}
	if index == 0 {
		anchorNextNode := l.firstNode
		l.firstNode = other.firstNode
		other.lastNode.nextNode = anchorNextNode
		l.length += other.length
		return true
	}
	if index == l.length {
		l.lastNode.nextNode = other.firstNode
		l.lastNode = other.lastNode
		l.length += other.length
		return true
	}
	anchorLeftNode := l.GetNodeByIndex(index - 1)
	anchorRightNode := l.GetNodeByIndex(index)
	anchorLeftNode.nextNode = other.firstNode
	other.lastNode.nextNode = anchorRightNode
	l.length += other.length
	return true
}

// ContainsNode проверяет наличие узла в списке
func (l *List) ContainsNode(node *Node) bool {
	if node == nil {
		return false
	}
	currentNode := l.firstNode
	for ; currentNode != l.lastNode.nextNode; currentNode = currentNode.nextNode {
		if currentNode == node {
			return true
		}
	}
	return false
}

// EqualValues проверяет равенство двух списков по значениям в узлах
func (l *List) EqualValues(other *List) bool {
	if l.length != other.length {
		return false
	}
	values, _ := l.GetAll()
	otherValues, _ := other.GetAll()
	if len(values) != len(otherValues) {
		return false
	}
	for i := 0; i < len(values); i++ {
		if values[i] != otherValues[i] {
			return false
		}
	}
	return true
}

// GetSubList возвращает подсписок
func (l *List) GetSubList(fromIndex, toIndex int64) *List {
	if fromIndex > toIndex || fromIndex < 0 || toIndex > l.Len() || toIndex < 0 {
		return nil
	}
	var currentNode *Node = l.firstNode
	var currentIndex int64 = 0
	for ; currentIndex < fromIndex; currentIndex++ {
		currentNode = currentNode.nextNode
	}
	subList := NewList()
	for ; currentIndex < toIndex; currentIndex++ {
		subList.Add(currentNode.value)
		currentNode = currentNode.nextNode
	}
	return subList
}
