/*The commented code is basically a simple way of claiming the files from the directory and should be followed for small projects*/

package renders

import (
	"github.com/SupraIZ/booking/packages/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/SupraIZ/booking/packages/config"
)

var app *config.AppConfig

//NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a // set global variable for use in templates
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

//Renders a Template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	
	//get the template cache from the app config --new

	var tc map[string]*template.Template

	if app.UseCache{
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing",err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.gohtml from ./templates
	pages, err := filepath.Glob("templates/*.page.gohtml")
	if err!=nil {
		return myCache, err
	}

	//range through all the files ending with *.page.gohtml
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache,err
		}
		matches, err := filepath.Glob("templates/*.layout.gohtml")
		if err!=nil {
			return myCache,err
		}

		if len(matches)> 0{
			ts, err = ts.ParseGlob("templates/*.layout.gohtml")
			if err!=nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache , nil
	// return the map and an error if any occurred while creating it or reading existing templates into memory
}