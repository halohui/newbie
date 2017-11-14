package main

import (
	"fmt"
	"newbie/hello"
	"newbie/arith"
)

type FunType func(int,float32)int //声明一个函数

func main() {
	fmt.Println("Hello world!")
	hello.Test()

	data:=[]int{89,34,2,3,8,11,89,233,5,8,12,3,6}
	fmt.Println(data)
	//arith.InsertSort(data)
	//arith.BinInsertSort(data)
	//arith.ShellSort(data)
	//arith.BubbleSortBig(data)
	//arith.BubbleSortSmall(data)
	//arith.QuickSort(data,0,len(data)-1)
	//arith.SelectSort(data)
	//arith.HeapSort(data)
	arith.MergeSort(data)
	fmt.Println(data)


}

func init() {
	fmt.Println("This init function in main")
}