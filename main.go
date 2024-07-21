package main

import (
  "net/http"
  "html/template"
  "embed"
  "log"
  "os"
)

//go:embed templates
var templateFiles embed.FS

//go:embed static
var static embed.FS

var templates = map[string]*template.Template {
  "index": template.Must(template.ParseFS(templateFiles, "templates/layouts/application.html", "templates/index.html")),
  "archive": template.Must(template.ParseFS(templateFiles, "templates/layouts/application.html", "templates/archive.html")),
}

func archiveHandler(writer http.ResponseWriter, request *http.Request) {
  err := templates["archive"].Execute(writer, nil)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
  }
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
  err := templates["index"].Execute(writer, nil)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.Handle("/static/", http.FileServer(http.FS(static)))

  http.HandleFunc("/archive", archiveHandler)
  http.HandleFunc("/", indexHandler)

  log.Println("Listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
