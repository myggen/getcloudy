package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {

	///we create a new router to expose our api
	//to our users
	handler := http.NewServeMux()

	handler.HandleFunc("/api/predictions", predictions)

	conn := fmt.Sprintf("0.0.0.0:%d", 9797)
	fmt.Printf("Starting server on %s\n", conn)
	err := http.ListenAndServe(conn, handler)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
func predictions(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("/lustre/store/project/metproduction/products/camrec/predictions.json")
	if err != nil {
		http.Error(w, err.Error(),
				http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(bytes))
}
