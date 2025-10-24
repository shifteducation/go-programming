package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	param := request.URL.Query().Get("param")

	_, _ = writer.Write([]byte(fmt.Sprintf("id: %s, param: %s", id, param)))
}

func templateHandler(w http.ResponseWriter, _ *http.Request) {
	data := ViewData{
		Title:   "World Cup",
		Message: "FIFA will never regret it",
	}
	p := path.Join("server", "templates", "index.html")
	tmpl, err := template.ParseFiles(p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("file not found"))
		return
	}
	_ = tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", rootHandler)
	mux.HandleFunc("GET /template", templateHandler)

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
