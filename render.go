package main

import (
	"errors"
	"html/template"
	"io"
	"os"
	"path"
	"sync"
)

type HTMLRender struct {
	templates   map[string]*template.Template
	templateDir string
}

var render *HTMLRender
var once sync.Once

func GetInstance() *HTMLRender {
	once.Do(func() {

		if TemplatesDir == "" {
			TemplatesDir = "templates/"
		}

		render = &HTMLRender{
			templates:   map[string]*template.Template{},
			templateDir: TemplatesDir,
		}
	})

	return render
}

func (r *HTMLRender) checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (r *HTMLRender) Render(wr io.Writer, templateName string, data interface{}) error {
	t, ok := r.templates[templateName]

	if ok {

		err := t.Execute(wr, data)
		return err

	} else {

		p := path.Join(r.templateDir, templateName)
		if r.checkFileIsExist(p) {
			parseFiles, err := template.ParseFiles(p)
			if err != nil {
				return errors.New(TemplateInitError)
			}
			r.templates[templateName] = parseFiles
			err = parseFiles.Execute(wr, data)
			return err
		} else {
			return errors.New(FileNotExistError)
		}
	}
}
