package add

import (
	_ "go_learn/day2/example4/test"
)

/*
当我们导入一个包，但是不引用里面的任何变量或函数时，只是做一个初始化时，可以在前面加_
example3中的方法还可以写成如下方式，先声明，赋值操作放在初始化函数（init）中。

每个源文件都可以包含一个init函数，这个init函数会在main函数执行之前自动被go运行框架调用

*/

func init() {
	Name = "hello world"
	Age = 1
}

var Name string = "hello"
var Age int = 100

/*

上面代码的执行顺序是
首先执行test包中的初始化函数，然后执行当前包中的变量声明并赋值
然后执行当前包中的init初始化函数，将Name和Age的值完成替换
最后执行main包的main函数
*/
