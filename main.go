package main

import (
	"fmt"
	"newbie/syntax"
	"unsafe"
	"math"
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
	//arith.RadixSort(data[:], 10)
	//fmt.Println(data)
	//syntax.DeclareArray()
	//syntax.DeclareSlice()
	syntax.AboutMap()
	//syntax.DeclareStruct()

	v:= struct {
		a byte
		b []int
		c byte
	}{}

	fmt.Println(unsafe.Alignof(v),unsafe.Sizeof(v))
	var t1 = [4]int{1,2,3,4}
	var t2=[...]int{1:67,9:234}
	fmt.Println(t1,t2)
	fmt.Println(2&1<<1)

	const pi1 =12
	fmt.Printf("%T\n", pi1) //float64
	const pi2 float64=math.Pi
	fmt.Printf("%T\n" ,pi2) //float64


	const p2 int32 =pi1 //为什么这个不可以？？
	const p3 float32 =pi1
	const p4 float64=pi1
	const p5 complex64 =pi1
	const p6 complex128=pi1

	var b1[0]int   //数组长度可以为0
	fmt.Println(b1)

}

/*func init() {
	fmt.Println("This init function in main!")
}*/

func Hello() {
	fmt.Println("hello world!")
}
