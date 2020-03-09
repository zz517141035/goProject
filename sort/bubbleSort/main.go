package main

import "fmt"

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前数组arr=", (*arr))

	//定义一个临时变量用于交换
	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp

			}
		}
	}

	fmt.Println("排序后数组arr=", (*arr))
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}

	BubbleSort(&arr)
}
