package bookcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/albar2305/go-book/entities"
	"github.com/albar2305/go-book/helper"
	"github.com/albar2305/go-book/models/bookmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	books := bookmodel.GetAll()
	data := map[string]any{
		"books": books,
	}

	temp, err := template.ParseFiles("views/book/index.html")
	helper.PanicIfError(err)

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/create.html")
		helper.PanicIfError(err)

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var book entities.Book

		publishYear, err := strconv.Atoi(r.FormValue("publish_year"))
		helper.PanicIfError(err)

		book.Title = r.FormValue("title")
		book.Author = r.FormValue("author")
		book.Publisher = r.FormValue("publisher")
		book.PublishYear = int64(publishYear)

		if ok := bookmodel.Create(book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		book := bookmodel.Detail(id)

		data := map[string]any{
			"book": book,
		}

		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var book entities.Book

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		publishYear, err := strconv.Atoi(r.FormValue("publish_year"))
		if err != nil {
			panic(err)
		}

		book.Title = r.FormValue("title")
		book.Author = r.FormValue("author")
		book.Publisher = r.FormValue("publisher")
		book.PublishYear = int64(publishYear)

		if ok := bookmodel.Edit(id, book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	if err := bookmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	book := bookmodel.Detail(id)
	data := map[string]any{
		"book": book,
	}

	temp, err := template.ParseFiles("views/book/detail.html")
	helper.PanicIfError(err)

	temp.Execute(w, data)
}
