package grtrack

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		if err != nil {
			ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	err = t.Execute(w, Info.Art)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Path[8:]
	indx, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	if indx > len(Info.Art) || indx < 1 {
		ErrorHandler(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("static/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	indx = indx - 1
	artistInfo := &ForArtistPage{
		Art: Info.Art[indx],
		Rel: Info.Rel.Index[indx],
	}

	err = t.Execute(w, artistInfo)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, status string, errorcase int) {
	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500: Internal server error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(errorcase)
		tmpl.Execute(w, status)
	}
}
