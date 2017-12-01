package syntax

import (
	"fmt"
)

func AboutMap() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals { //只使用key,因此不需要_
		fmt.Println(key, capitals[key])
	}

	var a1 map[int]string  //空map，只读不写
	a1 = map[int]string{5: "chen"}
	fmt.Println(a1)


	var a2 = make(map[[5]int]string) //数组也可以为key，只要支持==,！=
	fmt.Println(a2)
	b1:=[5]int{23,6,9,10}
	a2[b1]="xiao"
	fmt.Println(a2)

	a3 := map[int]string{1: "xiao", 63: "chen", 45: "hui"}
	fmt.Println(a3)

	for key :=range a3 { //key是随机的排序
		fmt.Print(key,"\t",a3[key])
	}
	fmt.Println()

	a3[1]+="xiao"  //可以直接赋值

	for key :=range a3 { //key是随机的排序
		fmt.Print(key,"\t",a3[key],"\t")
	}
	fmt.Println()

	a4:=map[int][4]int{1:{2,3}} //value是数组的初始化
	fmt.Println(a4)

	/*a4[1][2]=34  //值是数组或者结构体时，不能直接赋值
	fmt.Println(a4)*/

}
