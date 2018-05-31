package syntax

import "fmt"

func DeclareStruct() {
	type node struct {
		int //可以没有名字，但是int类型只有一个可以没有名字
		id   int
		next *node
	}
	n1 := node{//使用命名初始化方式，即部分初始化
		id: 1, //逗号不能省略
	}

	fmt.Println(n1)

	n2 := node{//部分初始化
		id: 2,
		next: &n1, //逗号不能省略
	}
	fmt.Println(n2)

	n3 := node{//全部初始化的时候应该要全部初始化，否则出错
		12,
		12,
		&n2,
	}

	fmt.Println(n3)

	person := struct { //定义匿名结构体变量
		name string
		age  byte
	}{
		name: "XH",
		age:  26,
	}

	fmt.Println(person)

	type file  struct {
		name string
		attr struct{ //匿名结构体成员变量
			owner int
			perm int
		}
	}

	f:=file{
		name:"XH",
		attr: struct {
			owner int
			perm  int
		}{owner:23 , perm: 23},
	}

	fmt.Println(f)

}
