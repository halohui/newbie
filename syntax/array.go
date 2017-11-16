package syntax

import (
	"fmt"
	"reflect"
)

func DeclareArray() {
	var d1 [3]int
	var d2 [3]int
	fmt.Println(d1 == d2) //数组可以直接比较，但是不能直接赋值
	//fmt.Println(d1==nil)  //不能和nil比较，因为类型不一致

	var d3 [4]int
	var d4 = [4]int{2, 3}         //var d4 [4]int{2,3}是不正确的
	var d5 = [4]int{1: 23, 3: 67} //索引初始化的方式
	fmt.Println(d3, d4, d5)

	var d6 = new([5]int) //d6的类型是*[5]int,这个可以联想C++中的new
	var d7 [5]int        //d6和d7的类型是不一样的
	fmt.Println(reflect.TypeOf(d6))
	fmt.Println(reflect.TypeOf(d7))

	var a1 = [...]int{1, 2, 3} //这样声明是数组
	var a2 = []int{3, 4, 5, 6} //这样声明是切片
	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(a2, reflect.TypeOf(a2))

	var a3 = [6]string{2: "chen", 4: "xiao"}   //索引初始化的方式
	var a4 = [...]string{1: "chen", 6: "xiao"} //索引初始化的方式
	fmt.Println(a3, reflect.TypeOf(a3))
	fmt.Println(a4, reflect.TypeOf(a4))

	a5 := [3]int{1, 2, 3} //数组
	a6 := []int{1, 2, 3}  //切片
	fmt.Println(a5, reflect.TypeOf(a5))
	fmt.Println(a6, reflect.TypeOf(a6))

	var x [4]int
	x=[4]int{3,4,6,7}  //先声明，后初始化
	fmt.Println(x)
}

func DeclareSlice() {
	var a1 = []int{23, 45, 89}     //第一种方式声明切片
	var a2 = []int{2: 100, 56: 99} //第二种方式声明
	var a3 = make([]int, 10)       //第3种方式
	var a4 = make([]int, 3, 10)    //第4种方式
	var a5 = new([10]int)[2:8]     //第5种方式

	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(a2, reflect.TypeOf(a2))
	fmt.Println(a3, reflect.TypeOf(a3))
	fmt.Println(a4, reflect.TypeOf(a4))
	fmt.Println(a5, reflect.TypeOf(a5))

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // 前闭后开区间，不包括最后一个索引的指向的元素

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4] //slice1被赋值之前的cap是4，因此end的值可以取4，这种方式也叫切片重组
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	c1:=[]int{1,2,3,6}
	c2:=make([]int,10)
	fmt.Println(c2)
	copy(c2,c1)  //切片复制
	fmt.Println(c2)
	c3 :=append(c1,7,9,1,90,107,11) //追加,会修改切片的容量
	fmt.Println(c3,len(c3),cap(c3))
	fmt.Println(len(c2),cap(c2))
	fmt.Println(len(c1),cap(c1))
}
