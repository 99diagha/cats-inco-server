package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		http.ServeFile(res, req, "./static/catdata.json")
	})

	http.ListenAndServe(":9000", nil)
}
