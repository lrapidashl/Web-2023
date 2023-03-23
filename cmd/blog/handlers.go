package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	MostRecent    []mostRecentData
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
	PostLink    string
}

type mostRecentData struct {
	Title       string
	Subtitle    string
	Img         string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Println(err.Error())
		return
	}

	data := indexPage{
		Title:         "Escape.",
		FeaturedPosts: featuredPosts(),
		MostRecent:    mostRecent(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
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

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "featured-post_image_the-road-ahead",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "September 25, 2015",
			PostLink:    "/post",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you’ve never been before.",
			ImgModifier: "featured-post_image_from-top-down",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "September 25, 2015",
			PostLink:    "#",
		},
	}
}

func mostRecent() []mostRecentData {
	return []mostRecentData{
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Img:         "static/img/still_standing_tall.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it’s going to be.",
			Img:         "static/img/sunny_side_up.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Img:         "static/img/water_falls.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Througt the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Img:         "static/img/through_the_mist.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Img:         "static/img/awaken_early.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Img:         "static/img/try_it_always.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
	}
}
