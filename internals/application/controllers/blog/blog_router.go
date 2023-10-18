package blog

import (
	"smux/internals/common/constants/routerconst"

	"github.com/go-chi/chi/v5"
)

func BlogRouter(router *chi.Mux, controller *BlogController) {
	router.Get(routerconst.Blogs, controller.GetBlogs)
}
