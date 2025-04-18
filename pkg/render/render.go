package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	cache, err := CreateCacheTemplate()
	if err != nil {
		log.Fatal(err)
	}

	if t, ok := cache[tmplt]; !ok {
		log.Fatalf("there is no template for the %s", tmplt)
	} else {
		buffer := new(bytes.Buffer)
		err = t.Execute(buffer, nil)
		if err != nil {
			fmt.Println("error parsing template %w", err)
		}

		_, err = buffer.WriteTo(w)
		if err != nil {
			fmt.Println("error to write to response writer from buffer %w", err)
		}
	}
}

func CreateCacheTemplate() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)
		parsedTmpl, err := template.New(filename).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			parsedTmpl, err = parsedTmpl.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[filename] = parsedTmpl
	}
	return cache, nil
}

/*var templateMap = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	var takeTemplate *template.Template
	var err error

	if _, ok := templateMap[tmplt]; !ok {
		err = createCacheTemplate(tmplt)
		if err != nil {
			fmt.Println("error creating cache template %w", err)
		}
	} else {
		fmt.Println("using template from cache")
	}

	takeTemplate = templateMap[tmplt]
	err = takeTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template %w", err)
	}
}

func createCacheTemplate(tmplt string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmplt),
		fmt.Sprintf("./templates/base.layout.tmpl"),
	}

	parsedTmpl, err := template.ParseFiles(templates...)
	if err != nil {
		fmt.Println("error parsing template %w", err)
		return err
	}

	templateMap[tmplt] = parsedTmpl
	return nil
}*/
