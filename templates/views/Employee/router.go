package Employee

import (
	"html/template"
	"net/http"

	"github.com/normanarauzl/goSimpleCrud/controllers"
)

var root = "/home/norman/go/src/github.com/normanarauzl/goSimpleCrud/templates/views/Employee/form/*"
var tmpl = template.Must(template.ParseGlob(root))

//var tmpl = template.Must(template.ParseGlob("/home/norman/go/src/github.com/normanarauzl/goSimpleCrud/templates/views/Employee/form/*"))

func GoIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", controllers.Index)
}

func GoEdit(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Edit", controllers.Edit)
}

func GoInsert(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Insert", controllers.Insert)
}

func GoShow(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Show", controllers.Show)
}

func GoNew(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}
