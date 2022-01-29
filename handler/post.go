package handler

import (
	"net/http"

	"github.com/afman42/go-crud-reactjs/helper"
	"github.com/afman42/go-crud-reactjs/post"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
	service post.Service
}

func NewPostHandler(service post.Service) *postHandler {
	return &postHandler{service}
}

func (h *postHandler) GetPosts(c *gin.Context) {

	posts, err := h.service.GetPosts()
	if err != nil {
		response := helper.APIResponse("Error to get posts", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of posts", http.StatusOK, "success", post.FormatPosts(posts))
	c.JSON(http.StatusOK, response)
}

func (h *postHandler) GetPostByID(c *gin.Context) {
	var input post.GetPostDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	postDetail, err := h.service.GetPostByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Post detail", http.StatusOK, "success", post.FormatPost(postDetail))
	c.JSON(http.StatusOK, response)
}

func (h *postHandler) CreatePost(c *gin.Context) {
	var input post.CreatePostInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create post", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPost, err := h.service.CreatePost(input)
	if err != nil {
		response := helper.APIResponse("Failed to create post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create post", http.StatusOK, "success", post.FormatPost(newPost))
	c.JSON(http.StatusOK, response)
}

func (h *postHandler) UpdatePost(c *gin.Context) {
	var inputID post.GetPostDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData post.CreatePostInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update post", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedPost, err := h.service.UpdatePost(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update post", http.StatusOK, "success", post.FormatPost(updatedPost))
	c.JSON(http.StatusOK, response)
}

func (h *postHandler) DeletePostByID(c *gin.Context) {
	var input post.GetPostDeleteInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get remove of post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, errr := h.service.DeletePostByID(input)
	if errr != nil {
		response := helper.APIResponse("Failed to get remove of post", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Post delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
