package main

import "fmt"

func InsertionSort(arr *[7]int) {

	for i := 0; i < len(arr)-1; i++ {
		insertVal := arr[i+1]
		insertIndex := i

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移操作
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入数据
		if insertIndex != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("完成第%d次操作结果为%v \n", i+1, *arr)
	}

	/*
		insertVal := arr[1]
		insertIndex := 1 - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移操作
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != 1 {
			arr[insertIndex+1] = insertVal
		}

		fmt.Println("完成第一次数据插入arr= ", *arr)

		//完成第二次操作 给第三个元素找到合适的位置
		insertVal = arr[2]
		insertIndex = 2 - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移操作
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != 2 {
			arr[insertIndex+1] = insertVal
		}

		fmt.Println("完成第二次数据插入arr= ", *arr)

		//完成第三次操作,给第四个元素找到合适位置
		insertVal = arr[3]
		insertIndex = 3 - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移操作
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != 3 {
			arr[insertIndex+1] = insertVal
		}

		fmt.Println("完成第三次数据插入arr= ", *arr)
	*/
}

func main() {
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	InsertionSort(&arr)
	fmt.Println(arr)
}
