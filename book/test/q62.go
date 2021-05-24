package main

// 将被测试的程序, 具体功能是取平均数
func funcDemo16(dataList []int) int {
	sum := 0
	for _, x := range dataList {
		sum += x
	}
	return sum / len(dataList)
}
