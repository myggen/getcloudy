package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	// Current reftime
	path := "/lustre/store/project/metproduction/products/camrec/predictions.json"

	reftimes, ok := r.URL.Query()["reftime"]

	if !ok || len(reftimes[0]) < 1 {
		log.Println("No parameter'reftime' given. Using current/latest reftime")

	} else {
		reftimeSstr := reftimes[0]
		reftime, err := time.Parse("20060102T1504Z", reftimeSstr)
		if err != nil {
			log.Printf("Error parsing reftime:  %v", err)
			w.WriteHeader(400)
			return
		}
		log.Printf("Reftime: %v", reftime)
		datadir := "/lustre/store/project/metproduction/products/camrec"
		path = fmt.Sprintf("%s/%s/%s/%s/predictions_%s.json", datadir,
			reftime.Format("2006"),
			reftime.Format("01"),
			reftime.Format("02"),
			reftime.Format("20060102T1504Z"),
		)

		log.Printf("Path: %s", path)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(bytes))
}
