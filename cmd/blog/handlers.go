package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	MostRecent    []mostRecentData
}

type featuredPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

type mostRecentData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Img         string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		mostRecentData, err := mostRecent(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			Title:         "Escape.",
			FeaturedPosts: featuredPostsData,
			MostRecent:    mostRecentData,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title, 
			subtitle,
			image_url,
			author, 
			author_url,
			publish_date 
		FROM
			post
		WHERE featured = 1
	`
	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func mostRecent(db *sqlx.DB) ([]mostRecentData, error) {
	const query = `
		SELECT
			title, 
			subtitle,
			image_url,
			author, 
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
	`
	var posts []mostRecentData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
