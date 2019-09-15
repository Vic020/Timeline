package hello

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

const (
	TemplateRootFolder = "templates"
)

type Hello struct {
	Name  string
	Path  string
	Param string
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["param"]

	if len(params) <= 0 {
		params = []string{""}
	}

	path_ := r.URL.Path

	hello := Hello{"Test", path_, params[0]}

	//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")

	fp := path.Join(TemplateRootFolder, "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(w, hello); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/foo", fooHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
