package datastruct

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: 601
 */

type Array struct {
	data   []interface{}
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]interface{}, capacity, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	// 底层是切片，所以不是大于length，而是大于cap才是越界
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (interface{}, error) {
	if this.isIndexOutOfRange(index) {
		return nil, errors.New("out of index range")
	}
	return this.data[index], nil
}

//插入数值到索引index上
func (this *Array) Insert(index uint, v interface{}) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	// 将index后的元素向后移动
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) InsertToTail(v interface{}) error {
	return this.Insert(this.Len(), v)
}

//删除索引index上的值
func (this *Array) Delete(index uint) (interface{}, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

//打印数列
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
