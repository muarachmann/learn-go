package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)



func handlerFucn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my portfolio</h1>")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handlerFucn)
	http.ListenAndServe(":3000", router)
}