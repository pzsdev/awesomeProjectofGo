package main

import "fmt"
import "net/http"

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// 简单的web服务器
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
