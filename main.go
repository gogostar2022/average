package main

import (
	"fmt"
)

func main() {
	//求数组的合
	numbers := [3] float64 {71.8,56.2,89.5}
	var sum float64= 0
	for _,number := range numbers{
		sum +=number
	}
	fmt.Println("sum:",sum)
	//求数组平均值
	size := float64(len(numbers))
	fmt.Println("average =",sum / size)
	
}