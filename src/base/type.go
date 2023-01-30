package main

import "testing"

type MyInt int64

func Test(t *testing.T)  {
	var  a int = 1
	var b int64 = 2
	// b = a  不支持隐式类型

	b = int64(a)
	t.Log(a ,b )
}
