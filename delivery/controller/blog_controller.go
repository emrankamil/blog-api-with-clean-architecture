package controller

import (
	"blog-api_with-clean-architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUseCase
}

func NewBlogController(bu domain.BlogUseCase) *BlogController {
	return &BlogController{
		BlogUsecase: bu,
	}
}


func (bc *BlogController) CreateBlog(cxt *gin.Context) {
	var blog domain.Blog
	err := cxt.ShouldBindJSON(&blog)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	createdBlog, err := bc.BlogUsecase.CreateBlog(cxt, &blog)
	// println(createdBlog, err.Error())

	if err != nil {
		cxt.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{"message": "Blog Created Successfully", "Blog": createdBlog})
	
}
func (bc *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")

	blog, err := bc.BlogUsecase.GetBlog(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, blog)
}



func (bc *BlogController) GetBlogs(c *gin.Context) {
	var pagination domain.Pagination
	err := c.BindQuery(&pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	blogs, err := bc.BlogUsecase.GetBlogs(c, &pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}



// update blog should be partially updated
func (bc *BlogController) UpdateBlog(cxt *gin.Context) {
	id := cxt.Param("id")
	var blog domain.Blog
	err := cxt.BindJSON(&blog)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}
	err = bc.BlogUsecase.UpdateBlog(cxt, &blog, id)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Updated Successfully"})
}

func (bc *BlogController) DeleteBlog(cxt *gin.Context) {
	id := cxt.Param("id")
	err := bc.BlogUsecase.DeleteBlog(cxt, id)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Deleted Successfully"})
}

func (bc *BlogController) LikeBlog(cxt *gin.Context) {
	blogID := cxt.Param("blog_id")
	userID := cxt.Param("user_id")
	err := bc.BlogUsecase.LikeBlog(cxt, blogID, userID)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Liked Successfully"})
}

func (ctrl *BlogController) UnlikeBlog(cxt *gin.Context) {
	blogID := cxt.Param("blogID")
	userID := cxt.Param("userID")

	err := ctrl.BlogUsecase.UnlikeBlog(cxt, blogID, userID)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

// CommentBlog handles HTTP requests to add a comment to a blog post
func (ctrl *BlogController) CommentBlog(cxt *gin.Context) {
	blogID := cxt.Param("blogID")

	var comment domain.Comment
	if err := cxt.BindJSON(&comment); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := ctrl.BlogUsecase.CommentBlog(cxt, blogID, &comment)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
