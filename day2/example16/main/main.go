package main

import "fmt"

/*

一般：
%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
%#v	相应值的Go语法表示
%T	相应值的类型的Go语法表示
%%	字面上的百分号，并非值的占位符

布尔：
%t	单词 true 或 false。

整数：
%b	二进制表示
%c	相应Unicode码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由Go语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode格式：U+1234，等同于 "U+%04X"

浮点数及其复合构成：
%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat
	的 'b' 转换格式一致。例如 -123456p-78
%e	科学计数法，例如 -1234.456e+78
%E	科学计数法，例如 -1234.456E+78
%f	有小数点而无指数，例如 123.456
%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出

字符串与字节切片：
%s	字符串或切片的无解译字节
%q	双引号围绕的字符串，由Go语法安全地转义
%x	十六进制，小写字母，每字节两个字符
%X	十六进制，大写字母，每字节两个字符

指针：
%p	十六进制表示，前缀 0x
*/
func main() {
	var a int
	var b bool
	c := 'a' // 单引号是bytes

	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("90%%\n")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100)
	fmt.Printf("%f\n", 199.22)
	fmt.Printf("%q\n", "this is a my_test")
	fmt.Printf("%x\n", 39839333)
	fmt.Printf("%p\n", &a)

	str := fmt.Sprintf("a=%d", a)
	fmt.Printf("%q\n", str)

	str1 := "hello"
	str2 := "world"

	//str3 := str1 + " " + str2
	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)
	fmt.Println(str3)
	fmt.Printf("len(str3)=%d\n", n)
	substr := str3[0:5]
	fmt.Println(substr)
	substr = str3[:5]
	fmt.Println(substr)
	substr = str3[6:]
	fmt.Println(substr)
}
