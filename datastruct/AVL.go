package datastruct

import "fmt"

type AVL struct {
	value  int  //值
	height int  //深度
	left   *AVL //左子树
	right  *AVL //右子树
}

//查找元素
func (t *AVL) Search(value int) bool {
	if t == nil {
		return false
	}
	compare := value - t.value
	if compare < 0 {
		return t.left.Search(value)
	} else if compare > 0 {
		return t.right.Search(value)
	} else {
		return true
	}
}

func (t *AVL) leftRotate() *AVL { //左旋转
	headNode := t.right
	t.right = headNode.left
	headNode.left = t
	//更新结点高度
	t.height = max(t.left.getHeight(), t.right.getHeight()) + 1
	headNode.height = max(headNode.left.getHeight(), headNode.right.getHeight()) + 1
	return headNode
}

func (t *AVL) rightRotate() *AVL { //右旋转
	headNode := t.left
	t.left = headNode.right
	headNode.right = t
	//更新结点高度
	t.height = max(t.left.getHeight(), t.right.getHeight()) + 1
	headNode.height = max(headNode.left.getHeight(), headNode.right.getHeight()) + 1
	return headNode
}

func (t *AVL) rightThenLeftRotate() *AVL { //右旋转,之后左旋转
	//以失衡点右结点先右旋转
	sonHeadNode := t.right.rightRotate()
	t.right = sonHeadNode
	//再以失衡点左旋转
	return t.leftRotate()
}

func (t *AVL) LeftThenRightRotate() *AVL { //左旋转,之后右旋转
	//以失衡点左结点先左旋转
	sonHeadNode := t.left.leftRotate()
	t.left = sonHeadNode
	//再以失衡点左旋转
	return t.rightRotate()
}

func (t *AVL) adjust() *AVL {
	if t.right.getHeight()-t.left.getHeight() == 2 {
		if t.right.right.getHeight() > t.right.left.getHeight() {
			t = t.leftRotate()
		} else {
			t = t.rightThenLeftRotate()
		}
	} else if t.left.getHeight()-t.right.getHeight() == 2 {
		if t.left.left.getHeight() > t.left.right.getHeight() {
			t = t.rightRotate()
		} else {
			t = t.LeftThenRightRotate()
		}
	}
	return t
}

//添加元素
func (t *AVL) Insert(value int) *AVL {
	if t == nil {
		newNode := AVL{value, 1, nil, nil}
		return &newNode
	}
	if value < t.value {
		t.left = t.left.Insert(value)
		t = t.adjust()
	} else if value > t.value {
		t.right = t.right.Insert(value)
		t = t.adjust()
	} else {
		fmt.Println("the node exit")
	}
	t.height = max(t.left.getHeight(), t.right.getHeight()) + 1
	return t
}

/*删除元素
*1、如果被删除结点只有一个子结点，就直接将A的子结点连至A的父结点上，并将A删除
*2、如果被删除结点有两个子结点，将该结点右子数内的最小结点取代A。
*3、查看是否平衡,该调整调整
 */
func (t *AVL) Delete(value int) *AVL {
	if t == nil {
		return t
	}
	compare := value - t.value
	if compare < 0 {
		t.left = t.left.Delete(value)
	} else if compare > 0 {
		t.right = t.right.Delete(value)
	} else { //找到结点,删除结点（）
		if t.left != nil && t.right != nil {
			t.value = t.right.getMin()
			t.right = t.right.Delete(t.value)
		} else if t.left != nil {
			t = t.left
		} else { //只有一个右孩子或没孩子
			t = t.right
		}
	}
	if t != nil {
		t.height = max(t.left.getHeight(), t.right.getHeight()) + 1
		t = t.adjust()
	}
	return t
}

//按顺序获得树中元素
func (t *AVL) getAll() []int {
	values := []int{}
	return addValues(values, t)
}

//将一个节点加入切片中
func addValues(values []int, t *AVL) []int {
	if t != nil {
		values = addValues(values, t.left)
		values = append(values, t.value)
		fmt.Println(t.value, t.height)
		values = addValues(values, t.right)
	}
	return values
}

//查找子树最小值
func (t *AVL) getMin() int {
	if t == nil {
		return -1
	}
	if t.left == nil {
		return t.value
	} else {
		return t.left.getMin()
	}
}

//查找子树最大值
func (t *AVL) getMax() int {
	if t == nil {
		return -1
	}
	if t.right == nil {
		return t.value
	} else {
		return t.right.getMax()
	}
}

//查找最小结点
func (t *AVL) getMinNode() *AVL {
	if t == nil {
		return nil
	} else {
		for t.left != nil {
			t = t.left
		}
	}
	return t
}

//查找最大结点
func (t *AVL) getMaxNode() *AVL {
	if t == nil {
		return nil
	} else {
		for t.right != nil {
			t = t.right
		}
	}
	return t
}

//得到树高
func (t *AVL) getHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
