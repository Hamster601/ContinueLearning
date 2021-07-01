package datastruct

// 链表的一个节点
type DoubleListNode struct {
	prev  *DoubleListNode // 前一个节点
	next  *DoubleListNode // 后一个节点
	value string          // 数据
}

// 创建一个节点
func NewDoubleListNode(value string) (doublelistNode *DoubleListNode) {
	doublelistNode = &DoubleListNode{
		value: value,
	}

	return
}

// 当前节点的前一个节点
func (n *DoubleListNode) Prev() (prev *DoubleListNode) {
	prev = n.prev

	return
}

// 当前节点的前一个节点
func (n *DoubleListNode) Next() (next *DoubleListNode) {
	next = n.next

	return
}

// 获取节点的值
func (n *DoubleListNode) GetValue() (value string) {
	if n == nil {

		return
	}
	value = n.value

	return
}

// 链表
type DoubleList struct {
	head *DoubleListNode // 表头节点
	tail *DoubleListNode // 表尾节点
	len  int             // 链表的长度
}

// 创建一个空链表
func NewDoubleList() (list *DoubleList) {
	list = &DoubleList{}
	return
}

// 返回链表头节点
func (l *DoubleList) Head() (head *DoubleListNode) {
	head = l.head

	return
}

// 返回链表尾节点
func (l *DoubleList) Tail() (tail *DoubleListNode) {
	tail = l.tail

	return
}

// 返回链表长度
func (l *DoubleList) Len() (len int) {
	len = l.len

	return
}

// 在链表的右边插入一个元素
func (l *DoubleList) RPush(value string) {

	node := NewDoubleListNode(value)

	// 链表未空的时候
	if l.Len() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.len = l.len + 1

	return
}

// 从链表左边取出一个节点
func (l *DoubleList) LPop() (node *DoubleListNode) {

	// 数据为空
	if l.len == 0 {

		return
	}

	node = l.head

	if node.next == nil {
		// 链表未空
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
	}
	l.len = l.len - 1

	return
}

// 通过索引查找节点
// 查不到节点则返回空
func (l *DoubleList) Index(index int) (node *DoubleListNode) {

	// 索引为负数则表尾开始查找
	if index < 0 {
		index = (-index) - 1
		node = l.tail
		for true {
			// 未找到
			if node == nil {

				return
			}

			// 查到数据
			if index == 0 {

				return
			}

			node = node.prev
			index--
		}
	} else {
		node = l.head
		for ; index > 0 && node != nil; index-- {
			node = node.next
		}
	}

	return
}

// 返回指定区间的元素
func (l *DoubleList) Range(start, stop int) (nodes []*DoubleListNode) {
	nodes = make([]*DoubleListNode, 0)

	// 转为自然数
	if start < 0 {
		start = l.len + start
		if start < 0 {
			start = 0
		}
	}

	if stop < 0 {
		stop = l.len + stop
		if stop < 0 {
			stop = 0
		}
	}

	// 区间个数
	rangeLen := stop - start + 1
	if rangeLen < 0 {

		return
	}

	startNode := l.Index(start)
	for i := 0; i < rangeLen; i++ {
		if startNode == nil {
			break
		}

		nodes = append(nodes, startNode)
		startNode = startNode.next
	}

	return
}
