package rest

import (
	"github.com/faruoqi/evento-events/controllers"
	"github.com/faruoqi/evento/model"
	"github.com/labstack/echo/v4"
)

func ServeAPI(endpoint string, dbHandler model.DbEventHandler) {

	controller := controllers.NewEventController(dbHandler)
	e := echo.New()
	eventRoute := e.Group("/events")
	eventRoute.GET("", controller.FindAll)
	eventRoute.GET("/:SearchCriteria/:search", controller.FindEvent)
	eventRoute.POST("", controller.AddEvent)
	e.Logger.Fatal(e.Start(endpoint))
}
