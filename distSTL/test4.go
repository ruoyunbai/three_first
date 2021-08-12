package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/model/", http.StripPrefix("/model/", http.FileServer(http.Dir("./model"))))
	http.Handle("/gltf/", http.StripPrefix("/gltf/", http.FileServer(http.Dir("./gltf"))))
	

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write([]byte(content))
}
