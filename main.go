package main

import "fmt"

func main(){
	var num = []int{1,9,8,5,6,7,2}
	sortNum := BubbleSort(num)
	fmt.Println(sortNum)

}

// 冒泡排序 o(n²)改良版
func BubbleSort(num []int)[]int{
	//o(n²)
	for i:=0;i<len(num)-1;i++{ // o(n)
		var change = false
		for j :=0;j<len(num)-1-i;j++{ //o(n)
			if num[j] < num[j+1]{
				num[j],num[j+1] = num[j+1],num[j]
				change = true
			}
		}
		if change == false{
			return num
		}
	}
	return num
}

// 选择排序
func SelectSort(num []int)[]int {
	return num
}