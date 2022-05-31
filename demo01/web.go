package demo01

import "net/http"

func Generate_Web() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	http.ListenAndServe(":8080", nil)
}
