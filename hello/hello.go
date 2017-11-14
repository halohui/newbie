package hello

import (
	"fmt"
	"math"
)

const a int = 67

var x int = 45

func init() {
	fmt.Println("x=", x, "This is an init function!")
}
func Test() {
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxFloat64)
}
