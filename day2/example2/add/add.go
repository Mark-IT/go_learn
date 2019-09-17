package add

var Name string = "hello world"
var Age int = 1

/*
包中的变量，如果想在其他包中被调用，其首字母必须是大写

如果想要给定义的变量赋初始值，则需要在像上面那样，声明时进行赋值，这样在编译时就能确定初始值

或者先定义变量，再在函数中进行初始化

错误写法 1

var Name string
var Age int

Name = "hello world "
Age = 10

因为go是编译型语言，必须在入口函数中执行代码

23和24行的代码没有在入口中执行，也不在函数中，所以编译时会报错



错误写法 2

Name:=  "hello world "
Age := 10

写法与上面错误写法1等价




*/
