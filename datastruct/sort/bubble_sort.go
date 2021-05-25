package sort

// 这是优化后的bubblesort
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 提前退出标志
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				//此次冒泡有数据交换
				flag = true
			}
		}
		// 如果没有交换数据，提前退出
		if !flag {
			break
		}
	}
}
