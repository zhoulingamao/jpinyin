package main

import (
	"fmt"
	"strings"

	"./jpinyin"
)

func main() {

	arrResult := jpinyin.GetInitiallist("a1重庆长壳大桥")
	strResult := strings.Join(arrResult, ",")
	if len(strResult) > 10 {
		strResult = strResult[:10]
	}

	fmt.Println(strResult)
}
