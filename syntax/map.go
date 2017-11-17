package syntax

import (
	"fmt"
)

func AboutMap() {
	capitals := map[string]string{"Xrance": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println(key, capitals[key])
	}

	var a1 map[int]string
	fmt.Println(a1)

	a1 = map[int]string{}
	fmt.Println(a1)



}
