package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	FeaturedPosts []*featuredPostData
	MostRecent    []*mostRecentData
}

type featuredPostData struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_modifier"`
	Img         string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}

type mostRecentData struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Img         string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}

type postData struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Img      string `db:"image_url"`
	Content  string `db:"content"`
}

type createPostRequest struct {
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Author        string `json:"author_name"`
	AuthorImgName string `json:"author_avatar"`
	AuthorImg     string `json:"author_avatar_file"`
	PublishDate   string `json:"publish_date"`
	ImgName       string `json:"page_image"`
	Img           string `json:"page_image_file"`
	Content       string `json:"content"`
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Check string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		mostRecentPostsData, err := mostRecent(db)
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
			MostRecent:    mostRecentPostsData,
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

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Internal post id", http.StatusForbidden)
			log.Println(err.Error())
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err.Error())
				return
			}

			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
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

	log.Println("Request completed successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html")
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

	log.Println("Request completed successfully")
}

func featuredPosts(db *sqlx.DB) ([]*featuredPostData, error) {
	const query = `
		SELECT
			post_id,
			title, 
			subtitle,
			image_modifier,
			image_url,
			author, 
			author_url,
			publish_date 
		FROM 
			post
		WHERE 
			featured = 1
	`
	var posts []*featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID
	}

	return posts, nil
}

func mostRecent(db *sqlx.DB) ([]*mostRecentData, error) {
	const query = `
		SELECT
			post_id,
			title, 
			subtitle,
			image_url,
			author, 
			author_url,
			publish_date
		FROM
			post
		WHERE 
			featured = 0
	`
	var posts []*mostRecentData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID
	}

	return posts, nil
}

func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			content,
			image_url
		FROM
		 	post
		WHERE
		  	post_id = ?
	`

	var post postData

	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	req.ImgName = "static/img/" + req.ImgName
	req.AuthorImgName = "static/img/" + req.AuthorImgName
	err := uploadImg(req.Img, req.ImgName)
	if err != nil {
		return err
	}
	err = uploadImg(req.AuthorImg, req.AuthorImgName)
	if err != nil {
		return err
	}

	const query = `
		INSERT INTO post
		(
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content
		)
		VALUES
		(
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
	`

	_, err = db.Exec(
		query,
		req.Title,
		req.Subtitle,
		req.Author,
		req.AuthorImgName,
		req.PublishDate,
		req.ImgName,
		req.Content,
	)
	return err
}

func uploadImg(imgFile string, imgName string) error {
	if imgName != "static/img/" {
		img, err := base64.StdEncoding.DecodeString(imgFile)
		if err != nil {
			return err
		}

		file, err := os.Create(imgName)
		if err != nil {
			return err
		}

		_, err = file.Write(img)
		if err != nil {
			return err
		}
	}
	return nil
}

func auth(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error1", 500)
			log.Println(err.Error())
			return
		}

		var req authRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal server error2", 500)
			log.Println(err.Error())
			return
		}

		check, err := checkUser(db)
		if err != nil {
			http.Error(w, "Internal server error3", 500)
			log.Println(err.Error())
			return
		}

		var resp authResponse

		if check {
			resp.Check = "yes"
		}
		if !check {
			resp.Check = "no"
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Internal server error4", 500)
			log.Println(err.Error())
			return
		}
	}
}

func checkUser(db *sqlx.DB) (bool, error) {
	const query = `
		SELECT
			email,
			pass
		FROM
		 	auth
		WHERE
		  	post_id = ?
	`

	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		log.Println(rows)
	}
	return true, nil
}
