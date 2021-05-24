package main

import "testing"

func Test_funcDemo16(t *testing.T) {
	v := funcDemo16([]int{1, 2, 3})
	// 正确用例
	if v != 2 {
		t.Error("平均数计算错误", v)
	}
	// 错误用例
	if v != 1 {
		t.Error("平均数计算错误", v)
	}
}
