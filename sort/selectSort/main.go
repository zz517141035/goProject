package main

import "fmt"

func SelectSort(arr *[6]int) {
	//先完成将diy8ige最大值和arr[0]交换
	//假设arr[0]就是最大值

	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		//2. 遍历后面 1---[len(arr) -1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("第%d次 %v\n  ", j+1, *arr)
	}

}
func main() {
	arr := [6]int{10, 34, 20, 100, 1000, 789}
	fmt.Println(arr)
	SelectSort(&arr)

}
