package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const authCookieName = "authBlog"

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

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func admin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authByCookie(db, w, r)
		if err != nil {
			return
		}

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
		err := authByCookie(db, w, r)
		if err != nil {
			return
		}

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

func logIn(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error1", 500)
			log.Println(err.Error())
			return
		}

		var req userRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal server error2", 500)
			log.Println(err.Error())
			return
		}

		userId, err := checkUser(db, req)
		if err != nil {
			http.Error(w, "Incorrect password or email", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    authCookieName,
			Value:   fmt.Sprint(userId),
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 1),
		})

		w.WriteHeader(200)
	}
}

func checkUser(db *sqlx.DB, req userRequest) (int, error) {
	const query = `
		SELECT
		    user_id
		FROM
			` + "`user`" +
		`WHERE
		  	email = ? 
			AND ` + "`password`" + ` = ?
	`

	userId := 0
	err := db.Get(&userId, query, req.Email, req.Password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func authByCookie(db *sqlx.DB, w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			log.Println(err)
			return err
		}
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return err
	}

	userIDStr := cookie.Value

	const query = `
	    SELECT
			EXISTS(
				SELECT
					1
				FROM
					` + "`user`" + `
				WHERE
		  			user_id = ?
				)
	`
	var exist bool
	err = db.Get(&exist, query, userIDStr)
	if err != nil {
		return err
	}
	if !exist {
		http.Error(w, "failed to find user with id "+userIDStr, http.StatusUnauthorized)
		return fmt.Errorf("failed to find user with id %s", userIDStr)
	}

	return nil
}

func logOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    authCookieName,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, -1),
	})

	w.WriteHeader(200)
}
