package main

import (
	"fmt"
	"strings"
)

/*
strings.HasPrefix(s string, prefix string) bool：判断字符串s是否以prefix开头 。
strings.HasSuffix(s string, suffix string) bool：判断字符串s是否以suffix结尾。

 */

func urlProcess(url string) string {	// 判断一个url是否以http://开头，如果不是，则加上http://。
	result := strings.HasPrefix(url, "http//")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {		// 判断一个路径是否以“/”结尾，如果不是，则加上/。
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	var (
		url  string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Println(url)
	fmt.Println(path)

}
