package routes

import (
	"net/http"
	"smux/di/container"
	"smux/internals/application/controllers/blog"
	"smux/internals/common/constants/routerconst"
	"smux/internals/domain/valobj"

	"smux/internals/common/utils"

	"github.com/go-chi/chi/v5"
)

func MountAll(root chi.Router, version valobj.AppVersion, controllerRegistry *container.ControllerRegistry) {
	router := chi.NewRouter()

	// all handlers here
	router.Get(routerconst.HealthCheck, func(w http.ResponseWriter, req *http.Request) {
		utils.ResponseWithJSONData(w, http.StatusOK, struct{}{})
	})

	blog.BlogRouter(router, controllerRegistry.BlogController)

	// end of routes

	// mount
	root.Mount(version.ToString(), router)
}
