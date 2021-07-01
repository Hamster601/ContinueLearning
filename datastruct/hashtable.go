package datastruct

import "fmt"

// 简易版的哈希表实现
const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

// 哈希函数
func hashFunction(i, size int) int {
	return (i % size)
}

// 插入值
func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

// 输出值
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
		}
		fmt.Println()
	}
}

// 查找值
func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}
