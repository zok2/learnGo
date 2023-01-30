package main

//
//{
//"1":"2",
//"2":"3",
//"3":"4",
//"5":"4"}
//4: 1 3 4 5
//key是员工的名称，value是直属上级的名称
//设计一个算法，告诉你员工名称后，输出他的所有下级名称

func find(f string, tMap map[string]string) []string {
	a := []string{}
	for k, v := range tMap {
		if v == f {
			a = append(a, k)
		}
	}
	return a
}

func recursion(f string, tMap map[string]string) []string {
	a := []string{}
	k := find(f, tMap)
	for _, v := range k {
		if v != "" {
			a = append(a, v)
			a = append(a, recursion(v, tMap)...)
			delete(tMap, v)
		}
	}
	return a
}

func test(f string, tMap map[string]string) {
	a := make(map[string][]string)
	for k, _ := range tMap {
		a[k] = append(a[k], find(k, tMap)...)
	}
	b := []string{}
	b = append(b, a[f]...)

}

func main() {
	var tMap = make(map[string]string)
	tMap["1"] = "2"
	tMap["2"] = "3"
	tMap["6"] = "3"
	tMap["4"] = ""
	tMap["3"] = "4"
	tMap["5"] = "4"
	test("4", tMap)

}
