package main

import "fmt"

func main() {
	// 手机型号
	mobileModel := "Oppo"

	switch mobileModel {
	case "XiaoMi", "HuaWei", "Oppo", "Vivo":
		fmt.Println("安卓")
	case "Iphone":
		fmt.Println("苹果")
	}

}
