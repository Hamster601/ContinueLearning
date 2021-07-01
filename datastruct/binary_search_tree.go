package datastruct


type Item interface {
}

// 节点
type TNode struct {
	Key int
	Value Item
	left *TNode
	right *TNode
}

// 接下来实现以下方法。
type Noder interface {
	Insert(key int, value Item)
	Min() *Item
	Max() *Item
	Search(key int) bool
	InOrderTraverse(f func(item Item))
	PreOrderTraverse(f func(item Item))
	PostOrderTraverse(f func(item Item))
	String()
}

// ItemBinarySearchTree the binary search tree of the Items
type ItemBinarySearchTree struct {
	Root *TNode
}

// insert element
func (bst *ItemBinarySearchTree) Insert(key int, value Item)  {
	n := &TNode{key, value, nil, nil}

	if bst.Root == nil{
		bst.Root = n
	}else {
		//
		insertNode(bst.Root, n)
	}
}
func insertNode(node, newNode *TNode)  {
	if newNode.Key < node.Key{
		// 左边插入
		if node.left == nil{
			node.left = newNode
		}else {
			insertNode(node.left, newNode)
		}
	}else {
		if node.right == nil{
			node.right = newNode
		}else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *ItemBinarySearchTree) Min() *Item {
	n := bst.Root
	
	if n == nil{
		return nil
	}
	
	for {
		if n.left == nil{
			return &n.Value
		}
		n = n.left
	}
}

// 最大值，一定在树的最右边
func (bst *ItemBinarySearchTree) Max() *Item  {
	n := bst.Root

	if n == nil{
		return &n.Value
	}

	for {
		if n.right == nil{
			return &n.Value
		}else {
			n = n.right
		}
	}
}

// 二叉搜索树，查找元素
func (bst *ItemBinarySearchTree) Search(key int) bool  {
	return search(bst.Root, key)
}

func search(n *TNode, key int) bool  {
	if n == nil{
		return false
	}

	if key < n.Key{
		return search(n.left, key)
	}

	if key > n.Key{
		return search(n.right, key)
	}

	return true
}

// 中序遍历
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item Item))  {
	inOrderTraver(bst.Root, f)
}

func inOrderTraver(n *TNode, f func(item Item))  {
	if n != nil{
		inOrderTraver(n.left, f)
		f(n.Value)
		inOrderTraver(n.right, f)
	}
}

// 前序遍历
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(item Item)) {
	preOrdertraver(bst.Root, f)
}

func preOrdertraver(n *TNode, f func(item Item))  {
	if n != nil{
		f(n.Value)
		preOrdertraver(n.left, f)
		preOrdertraver(n.right, f)
	}
}

// 后序遍历
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(item Item))  {
	postOrderTraver(bst.Root, f)
}

func postOrderTraver(n *TNode, f func(item Item))  {
	if n != nil{
		postOrderTraver(n.left, f)
		postOrderTraver(n.right, f)
		f(n.Value)
	}
}