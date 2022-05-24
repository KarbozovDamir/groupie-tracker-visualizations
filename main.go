package main

import (
	"fmt"

	"net/http"

	gr "grtrack/structure"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	gr.GetApi()
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", gr.MainPage)
	http.HandleFunc("/artist/", gr.ArtistPage)
	fmt.Print("listening to port: http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
