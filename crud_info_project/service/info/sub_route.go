package info

import (
	"info_project_service/service/info/controller"

	"github.com/go-chi/chi"
)

var ProjectInfoServiceSunRouter = chi.NewRouter()

func init() {
	ProjectInfoServiceSunRouter.Group(func(r chi.Router) {
		ProjectInfoServiceSunRouter.Post("/add", controller.AddProjectInfo)
		ProjectInfoServiceSunRouter.Get("/all", controller.AddProjectInfo)
		ProjectInfoServiceSunRouter.Patch("/update/id={id}", controller.UpdateProjectInfo)
		ProjectInfoServiceSunRouter.Delete("/remove/id={id}", controller.RemoveProject)

		ProjectInfoServiceSunRouter.With(auth.JWT).Post()
	})

}
