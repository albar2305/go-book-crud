package homecontroler

import (
	"html/template"
	"net/http"

	"github.com/albar2305/go-book/helper"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	helper.PanicIfError(err)

	temp.Execute(w, nil)
}
