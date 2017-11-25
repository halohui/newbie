package syntax

import (
	"fmt"
	"reflect"
)

/****************************************************************************
* 功能描述: 数组定义和初始化的方式
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-20 19:33:10          0.00           cxh                创建
*
*****************************************************************************/
func DeclareArray() {
	var d1 [3]int //声明数组变量，数组中元素全部被初始化为0
	var d2 [3]int = [3]int{12, 23, 45}
	fmt.Println(d1 == d2) //数组可以直接比较，同一类型的数组可以赋值，不同类型的数组不支持赋值，但可以比较
	//fmt.Println(d1==nil)  //不能和nil比较，因为类型不一致

	var d3 [4]int
	var d4 = [4]int{2, 3}         //var d4 [4]int{2,3}是不正确的
	var d5 = [4]int{1: 23, 3: 67} //索引初始化的方式,如果前面给出了数组的大小，编译器会进行索引检查

	fmt.Println(d3, d4, d5)

	var d6 = new([5]int) //d6的类型是*[5]int,这个可以和C++中的new类似
	var d7 [5]int        //d6和d7的类型是不一样的
	fmt.Println(reflect.TypeOf(d6))
	fmt.Println(reflect.TypeOf(d7))

	var a1 = [...]int{1, 2, 3} //这样声明是数组
	var a2 = []int{3, 4, 5, 6} //这样声明是切片，即[]中没有占位符的声明的就是切片
	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(a2, reflect.TypeOf(a2))

	var a3 = [6]string{2: "chen", 4: "xiao"}   //索引初始化的方式
	var a4 = [...]string{1: "chen", 6: "xiao"} //索引初始化的方式，这是数组
	fmt.Println(a3, reflect.TypeOf(a3))
	fmt.Println(a4, reflect.TypeOf(a4))

	a5 := [3]int{1, 2, 3} //数组
	a6 := []int{1, 2, 3}  //切片
	fmt.Println(a5, reflect.TypeOf(a5))
	fmt.Println(a6, reflect.TypeOf(a6))

	var a7 [4]int
	a7 = [4]int{3, 4, 6, 7} //先声明，后赋值初始化
	fmt.Println(a7)
	a8 := new(int)
	fmt.Println(a8)

	var b1 [5] int
	b1 = [5]int{2: 5, 4: 78} //这种赋值的方式也是支持的

	fmt.Println(b1)

	var data = [5]int{1, 2, 3, 4, 5}
	fmt.Println("data=", data)
	modifyArray(data)
	fmt.Println("data=", data)

	getFibonacci(50)
}

/****************************************************************************
* 功能描述: 切片的定义和初始化方式
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-20 19:33:03          0.00           cxh                创建
*
*****************************************************************************/
func DeclareSlice() {
	var a1 = []int{23, 45, 89}     //第1种方式声明切片，直接初始化底层数组
	var a2 = []int{2: 100, 56: 99} //第2种方式声明，使用索引初始化底层数组
	var a3 = make([]int, 10)       //第3种方式
	var a4 = make([]int, 3, 10)    //第4种方式
	var a5 = new([10]int)[2:8]     //第5种方式，使用new生成数组时取切片

	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(a2, reflect.TypeOf(a2))
	fmt.Println(a3, reflect.TypeOf(a3))
	fmt.Println(a4, reflect.TypeOf(a4))
	fmt.Println(a5, reflect.TypeOf(a5))

	var base [6]int
	var b1 []int = base[2:3:5] // 前闭后开区间，不包括最后一个索引的指向的元素

	for i := 0; i < len(base); i++ {
		base[i] = i
	}

	for i := 0; i < len(b1); i++ {
		fmt.Printf("%d\t", b1[i])
	}
	fmt.Println()

	fmt.Printf("The length of base is %d\n", len(base))
	fmt.Printf("The length of b1 is %d\n", len(b1))
	fmt.Printf("The capacity of b1 is %d\n", cap(b1))

	//slice1被赋值之前的cap是3，因此end的值可以取3，这种方式也叫切片重组
	//因此，end的值页不能够超出b1原来的cap3，否则会出现越界错误
	b1 = b1[0:3]
	for i := 0; i < len(b1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, b1[i])
	}

	fmt.Printf("The length of b1 is %d\n", len(b1))
	fmt.Printf("The capacity of b1 is %d\n", cap(b1))

	c1 := []int{1, 2, 3, 6, 8, 10, 11, 23, 26, 9, 10, 3, 4, 8, 9}
	c2 := make([]int, 3, 10)
	fmt.Println(c2)
	copy(c2, c1) //切片复制，复制内容以len小的为准
	fmt.Println(c2)
	c3 := append(c1, 7, 9, 1, 90, 107, 11)
	//追加,会修改切片的容量,即cap的值
	fmt.Println(c3, len(c3), cap(c3))
	fmt.Println(len(c2), cap(c2))
	fmt.Println(len(c1), cap(c1))

}

/****************************************************************************
* 功能描述: 验证go中数组是按值传递的
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 在函数中修改数组元素的值不会影响

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-25 21:22:40          0.00           cxh                创建
*
*****************************************************************************/
func modifyArray(data [5]int) {

	for i := 0; i < len(data); i++ {
		data[i] *= 10
	}

	fmt.Println("After modified in function data is:", data) //打印拷贝后的数组
}

/****************************************************************************
* 功能描述: 求Fibonacci数组元素的值
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: Fibonacci的第一项为

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-25 21:24:47          0.00           cxh                创建
*
*****************************************************************************/
func getFibonacci(len int) int {
	first, second, third := 1, 2, 3

	fmt.Print(first, "\t", second, "\t")
	for i := 2; i < len; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		third = first + second
		first = second
		second = third
		fmt.Print(third, "\t")
	}

	return third
}
