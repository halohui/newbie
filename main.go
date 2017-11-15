package main

import (
	"fmt"
	"newbie/arith"
	"newbie/syntax"
)

func main() {

	data := []int{89, 34, 2, 3, 8, 11, 89, 233, 5, 8, 12, 3, 6}
	//data:=[]int{9,3,2,3,8,1,8,3,5,8,2,3,6}
	fmt.Println(data)
	//arith.InsertSort(data)
	//arith.BinInsertSort(data)
	//arith.ShellSort(data)
	//arith.BubbleSortBig(data)
	//arith.BubbleSortSmall(data)
	//arith.QuickSort(data,0,len(data)-1)
	//arith.SelectSort(data)
	//arith.HeapSort(data)
	//arith.MergeSort(data)
	//arith.CountSort(data,10)
	arith.RadixSort(data[:], 10)
	fmt.Println(data)
	//syntax.DeclareArray()
	syntax.DeclareSlice()

}

func init() {
	fmt.Println("This init function in main!")
}
