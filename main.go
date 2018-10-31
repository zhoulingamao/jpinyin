package main

import (
	"fmt"
	"strings"

	"./jpinyin"
)

func main() {

	//调用GetInitiallist函数获取字符串拼音首字母，返回结果为[]string
	arrResult := jpinyin.GetInitiallist("a1重庆长壳大桥")

	strResult := strings.Join(arrResult, ",")

	fmt.Println(strResult)
}
