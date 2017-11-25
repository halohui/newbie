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

	a3 := map[int]string{1: "xiao", 63: "chen", 45: "test"}
	fmt.Println(a3)

	for key :=range a3 {
		fmt.Print(key,"\t")
	}

}
