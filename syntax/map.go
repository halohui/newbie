package syntax

import (
	"fmt"
)

func AboutMap() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println(key, capitals[key])
	}

}
