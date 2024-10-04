package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	indexTemp, err := template.ParseFiles("./web/dist/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	fileServer := http.FileServer(http.Dir("./web/dist"))

	http.Handle("/dist/", http.StripPrefix("/dist/", fileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = indexTemp.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})

	port := "3000"
	log.Printf("The client server is running on the port %v", port)
	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
