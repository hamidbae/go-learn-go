package controllers

import (
	"jwt-medium/models"
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

	p := models.Post{}
	p.Title = input.Title
	p.Description = input.Description
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
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "param not a number"})
	}

	p, err := models.GetOnePostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	err = p.DeletePost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"delete post success"})
}