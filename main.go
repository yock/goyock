package main

import (
  "net/http"
  "html/template"
  "embed"
  "log"
  "os"
  "fmt"
  "strings"
  "github.com/russross/blackfriday/v2"
  "github.com/adrg/frontmatter"
)

//go:embed templates
var templateFiles embed.FS

//go:embed static
var static embed.FS

//go:embed content
var content embed.FS

type Index struct {
  Title string
  Body []byte
}

type Frontmatter struct {
  Title string `yaml:"title"`
  PublishDate string `yaml:"publishDate"`
  Path string `yaml:"path"`
}

func renderMarkdown(args ...interface{}) template.HTML {
  s := blackfriday.Run([]byte(fmt.Sprintf("%s", args...)))
  return template.HTML(s)
}

var templates = map[string]*template.Template {
  "index": template.Must(template.New("index.html").Funcs(template.FuncMap{ "markdown": renderMarkdown }).ParseFiles("templates/index.html", "templates/layouts/application.html")),
  "archive": template.Must(template.ParseFS(templateFiles, "templates/layouts/application.html", "templates/archive.html")),
}


func archiveHandler(writer http.ResponseWriter, request *http.Request) {
  err := templates["archive"].Execute(writer, nil)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
  }
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
  file, err := content.Open("content/the-costliness-of-change/post.md")
  if err != nil {
    log.Fatal("Could not read file")
  }

  stat, err := file.Stat()
  if err != nil {
    log.Fatal("Could not stat file")
  }

  contents := make([]byte, stat.Size())
  _, err = file.Read(contents)
  if err != nil {
    log.Fatal(err)
  }

  var matter Frontmatter

  body, err := frontmatter.Parse(strings.NewReader(string(contents)), &matter)

  indexData := Index {
    Title: matter.Title,
    Body: body,
  }

  err = templates["index"].ExecuteTemplate(writer, "application", indexData)
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
