package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		url  string
		path string
	)
	// fmt.Println("input web address:")
	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)

	// fmt.Scanf("%s", &path)
	path = pathProcess(path)
	fmt.Println("url: ", url)
	fmt.Println("path: ", path)
}

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}
