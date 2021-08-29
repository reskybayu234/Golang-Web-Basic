// the function of http.HandleFunc() is usually used for route/endpoint registration and that handler.

package main

import "fmt"
import "net/http"

func main(){
	handlerIndex := func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello"))
	}

	http.HandleFunc("/",handlerIndex)
	http.HandleFunc("/index",handlerIndex)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Again"))
	})

	// menjalankan server
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}