package test

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Resty struct {
	Client *resty.Client
}

func NewResty() *Resty {
	return &Resty{
		Client: resty.New(),
	}
}

func (r *Resty) GetData() {
	var posts []Post

	res, err := r.Client.R().
		SetResult(&posts).
		Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatalf("Error fetching API: %v", err)
	}

	fmt.Println("Response Status Code:", res.Status())
	for _, post := range posts {
		fmt.Printf("Post ID: %d, Title: %s\n", post.ID, post.Title)
	}

}

func (r *Resty) GetDataWithPathParam(id int) {
	var post Post

	res, err := r.Client.R().
		SetResult(&post).
		Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))

	if err != nil {
		log.Fatalf("Error fetching API: %v", err)
	}

	fmt.Println("Response Status Code:", res.Status())
	fmt.Printf("Post ID: %d, Title: %s\n", post.ID, post.Title)
}

func (r *Resty) CreatePost(newPost Post) {
	var post Post
	res, err := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(newPost). // default request content type is JSON
		SetResult(&post). // or SetResult(LoginResponse{}).
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatalf("Error saat melakukan request: %v", err)
	}

	fmt.Println("Response Status:", res.Status())
	fmt.Printf("Hasil POST: %+v", post)
}

func (r *Resty) UpdatePost(id int, updatedPost Post) {
	var post Post
	res, err := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(updatedPost). // default request content type is JSON
		SetResult(&post).     // or SetResult(LoginResponse{}).
		Put(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))

	if err != nil {
		log.Fatalf("Error saat melakukan request: %v", err)
	}

	fmt.Println("Response Status:", res.Status())
	fmt.Printf("PUT Result: %+v", post)
}
