package syntax

import (
	"fmt"
)

func AboutMap() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println(key, capitals[key])
	}

	var a1 map[int]string
	a1 = map[int]string{5: "chen"}
	fmt.Println(a1)

	var x = make(map[[5]int]string)

	fmt.Println(x)

	t := map[int]string{1: "xiao", 63: "chen", 45: "test"}
	fmt.Println(t)

	for key :=range t{
		fmt.Print(key,"\t")
	}

}
