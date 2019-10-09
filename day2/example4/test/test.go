package test

import "fmt"

var Name string = "this is  in my_test package"
var Age int = 1000

func init() {
	fmt.Println("this is a my_test")
	fmt.Println("my_test.package.Name=", Name)
	fmt.Println("my_test.package.Age=", Age)
	Age = 10
	fmt.Println("my_test.package.Age=", Age)

}
