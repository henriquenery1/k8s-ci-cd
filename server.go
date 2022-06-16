package main

import "net/http"

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":80", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world! K8s</h1>"))
}
