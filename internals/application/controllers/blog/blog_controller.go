package blog

import (
	"net/http"
	"smux/internals/common/utils"
	"smux/internals/domain/usecases"
)

type IBlogController interface {
	GetBlogs(http.ResponseWriter, *http.Request)
}

type BlogController struct {
	blogUseCase *usecases.BlogUseCase
}

func NewBlogController(blogUseCase *usecases.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
	}
}

func (bc *BlogController) GetBlogs(w http.ResponseWriter, req *http.Request) {
	utils.ResponseWithJSONData(w, http.StatusOK, bc.blogUseCase.GetBlogs())
}
