package main

import (
	"fmt"
	"os"
	"text/template"
)

/*
2）if常见操作符

if常见操作符

	not 非{{if not .condition}} {{end}}
	and 与{{if and .condition1 .condition2}} {{end}}
	or 或{{if or .condition1 .condition2}} {{end}}
	eq 等于{{if eq .var1 .var2}} {{end}}
	ne 不等于{{if ne .var1 .var2}} {{end}}
	lt 小于 (less than){{if lt .var1 .var2}} {{end}}
	le 小于等于{{if le .var1 .var2}} {{end}}
	gt 大于{{if gt .var1 .var2}} {{end}}
	ge 大于等于{{if ge .var1 .var2}} {{end}}

循环


{{range.}}

{{end }}

 */

type Person struct {
	Name  string
	Title string
	Age   int
}

func main() {
	t, err := template.ParseFiles("E:/go_projects/src/go_learn/day10/template/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	//p :=map[string]string{"Name": "Mary", "Age": "31", "Title": "我的个人网站"}
	p := Person{Name: "Mary1", Age: 18, Title: "我的个人网站"}	// 必须都是大写才能访问

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
