package main

import "fmt"
const pi = 3.1415

//批量生成常量 is_admin is_delete 默认user_status

const (
	user_status = true
	is_admin
	is_delete
)
const http_default = 200
//iota 枚举 0 换行自增
const http_service_error = 500
const (
	http_200_code = iota + http_default
	http_201_code = iota + http_default
	http_202_code = iota + http_default
	http_203_code = iota + http_default
	http_204_code = iota + http_default
)
//iota从新开始计算
const http_500_code = iota+http_service_error

const (
	http_300_code = iota + http_default +100
	http_301_code = iota + http_default +100
	http_302_code = iota + http_default +100
	_ = iota + http_default +100
	http_304_code = iota + http_default +100
)
const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main()  {
	fmt.Println(pi)
	fmt.Println(user_status,is_admin,is_delete)
	fmt.Println(http_200_code,http_201_code,http_304_code,http_500_code)
	fmt.Println(KiB,MiB,GiB)
}