package landing

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("views/layouts/default.tmpl", "views/pages/home.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, map[string]interface{}{
		"message": "just coding...",
	})
}
