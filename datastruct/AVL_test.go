package datastruct

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	bsTree := AVL{100, 1, nil, nil}
	newTree := bsTree.Insert(60)
	newTree = bsTree.Insert(120)
	newTree = bsTree.Insert(110)
	newTree = bsTree.Insert(130)
	newTree = bsTree.Insert(105)
	fmt.Println(newTree.getAll())

	newTree.Delete(110)
	fmt.Println(newTree.getAll())

}
