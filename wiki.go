package main

import (
    "regexp"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

var dataPath="data/"
var templatePath="static/"

// template caching
// `template.Must` is a convenience wrapper that 
// panics when passed a non-nil error value, and 
// otherwise returns the *Tempalte unaltered.
var templates = template.Must(template.ParseFiles(
    templatePath + "edit.html", 
    templatePath + "view.html", 
    templatePath + "index.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := dataPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := dataPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        // err.Error sends a specified http response code 
        // in this case "Internal Server Error"
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
    if err != nil {
        // `http.Redirect` adds an HTTP status code of `http.StatusFound`(302)
        // and a Location header to the HTTP response.
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{ Title: title, Body: []byte(body) }
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index", nil)
}

func main() {
    http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
