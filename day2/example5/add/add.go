package add

/*
example3中的方法还可以写成如下方式，先声明，赋值操作放在初始化函数（init）中。

init函数中的代码，会在包导入时自动执行

*/

var Name string
var Age int

func init() {
	Name = "hello world"
	Age = 10
}
