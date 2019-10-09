package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "hello, world")
	//io.WriteString(w, "hello, world")
	panic("my_test my_test")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		fmt.Println(request.Form)
		request.ParseForm() // 执行后，Form中才能拿到值
		fmt.Println(request.Form)

		io.WriteString(w, request.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}
func main() {
	http.HandleFunc("/a", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc { // 自定义异常日志打印
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
