package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathToUrls := map[string]string {
		"/url-godoc" : "https://courses.calhoun.io/courses/cor_gophercises",
		"yaml-godoc" : "https://www.udemy.com/",
	}
	mapHandler := MapHandler(pathToUrls,mux)

	//Build the YAMLHandler using the maqHandler as the 
	// fallback 

	yaml := `
- path: /url-godoc
  url: https://courses.calhoun.io/courses/cor_gophercises
- path: /yaml-godoc
  url: https://www.udemy.com/
`

	yamlHandler ,err := YAMLHandler([]byte(yaml),mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("starting the server at port")
	http.ListenAndServe(":8080",yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/",hello)
	return mux
}

func hello(w http.ResponseWriter,req *http.Request) {
	fmt.Fprintln(w,"Hello , World ")
}