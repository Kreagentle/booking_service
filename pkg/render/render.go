package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmplt string) {
	parsedTmpl, _ := template.ParseFiles("./templates"+tmplt, "./templates/base.layout.tmpl")
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template %w", err)
	}
}

var templateMap = make(map[string]*template.Template)

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
}
