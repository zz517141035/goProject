package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(left int, right int, arr *[80000000]int) {
	l := left
	r := right
	//piovt是中轴 支点
	pivot := arr[(left+right)/2]

	//定义一个变量 辅助工作

	temp := 0

	//进行for循环 将比 pivot小的值放到左边
	//将比pivot大的值放到右边
	for l < r {
		//从pivot的左边找到大约等于pivot的值
		for arr[l] < pivot {
			l++
		}
		//从pivot右边知道到小于pivot的值
		for arr[r] > pivot {
			r--
		}
		//当l>=时表示本次任务完成 跳出
		if l >= r {
			break
		}
		//将值进行交换
		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		//考虑中间轴的值可能相等
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	//arr := [9]int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	var arr [80000000]int
	for i := 0; i < 80000000; i++ {
		arr[i] = rand.Intn(80000000)
	}

	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	//fmt.Println(arr)
	fmt.Printf("快速排序法耗时%d秒", end-start)

}
