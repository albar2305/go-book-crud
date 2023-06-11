package bookmodel

import (
	"github.com/albar2305/go-book/config"
	"github.com/albar2305/go-book/entities"
	"github.com/albar2305/go-book/helper"
)

func GetAll() []entities.Book {
	rows, err := config.DB.Query(`SELECT * FROM books`)
	helper.PanicIfError(err)

	defer rows.Close()

	var books []entities.Book

	for rows.Next() {
		var book entities.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.PublishYear)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

func Create(book entities.Book) bool {
	result, err := config.DB.Exec(`
		INSERT INTO books(
			title, author, publisher, publish_year
		) VALUES (?, ?, ?, ?)`,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishYear,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Book {
	row := config.DB.QueryRow(`
		SELECT 
			books.id, 
			books.title, 
			books.author, 
			books.publisher, 
			books.publish_year FROM books
		WHERE books.id = ?
	`, id)

	var book entities.Book

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.PublishYear,
	)

	if err != nil {
		panic(err)
	}

	return book
}

func Edit(id int, book entities.Book) bool {
	query, err := config.DB.Exec(`
		UPDATE books SET 
			title = ?, 
			author = ?,
			publisher = ?,
			publish_year = ?
		WHERE id = ?`,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishYear,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
