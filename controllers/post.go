package controllers

import (
	"jwt-medium/models"
	"jwt-medium/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPost(c *gin.Context) {
	posts, err := models.GetPosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":posts})
}

func GetOnePostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "param not a number"})
	}

	post, err := models.GetOnePostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":post})
}

type AddPostInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddPost(c *gin.Context){
	var input AddPostInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Post{}
	p.Title = input.Title
	p.Description = input.Description
	p.UserId = user_id
	_, err = p.SavePost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"add post success", "data":p})

}

type UpdatePostInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func UpdatePost(c *gin.Context){
	post_id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "param not a number"})
	}

	p, err := models.GetOnePostById(post_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.UserId != user_id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorized"})
		return
	}

	var input UpdatePostInput
	err = c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.Title = input.Title
	p.Description = input.Description
	_, err = p.UpdatePost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"update post success", "data":p})

}

func DeletePost(c *gin.Context){
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "param not a number"})
	}

	p, err := models.GetOnePostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.UserId != user_id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorized"})
		return
	}

	err = p.DeletePost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"delete post success"})
}

func LikePost(c *gin.Context){
	postId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "param not a number"})
	}

	_, err = models.GetOnePostById(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.LikePost(postId, int(user_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"like post success"})
}