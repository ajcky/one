package main

import "fmt"

func main() {
	var num = []int{1, 9, 8, 5, 6, 7, 2}
	//sortNum := BubbleSort(num)
	//fmt.Println(sortNum)
	//selectNum :=SelectSort(num)
	//fmt.Println(selectNum)

	//insertNum := InsertSort(num)
	//fmt.Println(insertNum)
	quickNum := QuickSort(num, 0, len(num)-1)
	fmt.Println(quickNum)
}

// 冒泡排序 o(n²)改良版
func BubbleSort(num []int) []int {
	//o(n²)
	for i := 0; i < len(num)-1; i++ { // o(n)
		var change = false
		for j := 0; j < len(num)-1-i; j++ { //o(n)
			if num[j] < num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				change = true
			}
		}
		if change == false {
			return num
		}
	}
	return num
}

// 选择排序
func SelectSort(num []int) []int {
	for i := 0; i < len(num)-1; i++ {
		min := i
		for j := i + 1; j < len(num); j++ {
			if num[j] < num[min] {
				min = j
			}
		}
		if min != i {
			num[min], num[i] = num[i], num[min]
		}
	}

	return num
}

// 插入排序
func InsertSort(num []int) []int {
	for i := 1; i < len(num); i++ {
		tmp := num[i]
		j := i - 1
		for {
			if num[j] > tmp && j >= 0 {
				num[j+1] = num[j]
				j--
			} else {
				break
			}
			num[j+1] = tmp
		}

	}
	return num
}
// 快速排序
func QuickSort(num []int, left, right int) []int {
	if left < right {
		mid := Partition(num, left, right)
		QuickSort(num, left, mid-1)
		QuickSort(num, mid+1, right)
	}
	return num
}
//
func Partition(num []int, left, right int) int {
	tmp := num[left]
	for left < right {
		for left < right && num[right] >= tmp {
			right --
		}
		num[left] = num[right]
		for left < right && num[left] <= tmp {
			left++
		}
		num[right] = num[left]
	}
	num[left] = tmp
	return left
}
