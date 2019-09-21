package main

import "fmt"

func testMap() {
	// map声明后，要使用，必须先初始化
	var a map[string]string // 声明是不会分配内存的，初始化需要make
	a = make(map[string]string, 10)

	// 或者声明的时候初始化
	//var a map[string]string = map[string]string{
	//	//"key": "value",
	//}

	a["abc"] = "efg"
	a["abc"] = "efg"
	a["abc1"] = "efg"

	fmt.Println(a)
}
func testMap2() {

	a := make(map[string]map[string]string) // key是字符串，val是map
	//a := make(map[string]map[string]string,100 )	// 如果知道大概容量，建议提前声明创建好，以免动态扩容
	a["key1"] = make(map[string]string) // 因为val是map，赋值时，需要提前初始化
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	fmt.Println(a)

}
func modify(a map[string]map[string]string) {
	val, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)

	} else {
		fmt.Println("val", val)
	}
	a["zhangsan"]["passwd"] = "123456"
	a["zhangsan"]["nickname"] = "pangpang"

	return
}
func testMap3() {

	a := make(map[string]map[string]string, 100)
	//a["zhangsan"] = map[string]string{
	//	"k1": "v1", // 必须有逗号
	//}
	modify(a)

	fmt.Println(a)
}

func trans(a map[string]map[string]string) {
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}
}

func testMap4() {

	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"

	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "abc"
	a["key2"]["key3"] = "abc"

	trans(a)          // 遍历
	delete(a, "key1") // 删除某个key

	// 如果要删除所有，则遍历删除，或者直接重新声明一下
	// 方法1
	//a = make(map[string]map[string]string)
	// 方法2
	//for k, _ := range a {
	//	delete(a, k)
	//}

	fmt.Println()
	trans(a)

	fmt.Println(len(a))
}

func testMap5() {
	var a []map[int]int
	a = make([]map[int]int, 5)

	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}

func main() {
	//testMap()
	//testMap2()
	//testMap3()
	testMap4()
	//testMap5()
}
