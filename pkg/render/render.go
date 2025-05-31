package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/models"
)

var functions = template.FuncMap{}

func AddDefaultData(data *models.TmpltData, r *http.Request) *models.TmpltData {
	data.CSRFToken = nosurf.Token(r)
	return data
}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmplt string, data *models.TmpltData) {
	var cache map[string]*template.Template
	var err error

	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, err = CreateCacheTemplate()
		if err != nil {
			fmt.Println("error create cache template %w", err)
		}
	}

	if t, ok := cache[tmplt]; !ok {
		log.Fatalf("there is no template for the %s", tmplt)
	} else {
		buffer := new(bytes.Buffer)

		data = AddDefaultData(data, r)
		err = t.Execute(buffer, data)
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
		log.Println("cant find page %w", err)
		return cache, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)
		parsedTmpl, err := template.New(filename).ParseFiles(page)
		if err != nil {
			log.Println("cant create new template %w", err)
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("cant parse layout %w", err)
			return cache, err
		}

		if len(layouts) > 0 {
			parsedTmpl, err = parsedTmpl.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("cant find a parsedTmpl %w", err)
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
