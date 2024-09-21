package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var posts = []Post{
	{ID: "1", Title: "jasim muhammed", Desc: "is Aloso A Boy"},
}

func getPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, posts)
}
func creatPost(ctx *gin.Context) {

	var newPost Post
	err := ctx.BindJSON(&newPost)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Enter Valid"})
		return
	}
	id := newPost.ID
	item, err := postById(id)
	if item != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This ID Alredy Used"})
		return
	}
	if err != nil {
		posts = append(posts, newPost)
		ctx.IndentedJSON(http.StatusCreated, newPost)
	}

}

func postById(id string) (*Post, error) {
	for i, v := range posts {
		if v.ID == id {
			return &posts[i], nil
		}
	}
	return nil, errors.New("Post Not Fount")
}

func updatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	item, err := postById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Post Not Font"})
		return
	}
	var newPost Post
	er := ctx.BindJSON(&newPost)
	if er != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Enter Valid Input"})
		return
	}

	item.Title = newPost.Title
	item.Desc = newPost.Desc
	ctx.IndentedJSON(http.StatusOK, item)
}

func deletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := postById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Post Not Font"})
		return
	}

	for index, val := range posts {
		if val.ID == id {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "post deleted"})

}

func getPostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	item, err := postById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Post Not Font"})
		return
	}
	ctx.JSON(http.StatusOK, item)

}
func main() {
	server := gin.Default()

	server.GET("/posts", getPost)
	server.POST("/post", creatPost)
	server.PATCH("/post/:id", updatePost)
	server.DELETE("/post/:id", deletePost)
	server.GET("/post/:id", getPostByID)

	server.Run()
}
