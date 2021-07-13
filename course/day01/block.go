package main

import "fmt"

var packageVar string = "aaaa"

func main() {
	var funcVar string = "bbbb"

	{ //限定变量的使用范围
		var blockVar string = "cccc"
		fmt.Println(packageVar, funcVar, blockVar)
	}

	fmt.Println(packageVar, funcVar)
}

//变量从自身的块开始读取，再往外读取，最后读取全局，还是没有就报错。
