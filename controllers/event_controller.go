package controllers

import (
	"fmt"
	"github.com/faruoqi/evento/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type eventServiceHandler struct {
	dbHandler model.DbEventHandler
}

func NewEventController(dbHandler model.DbEventHandler) *eventServiceHandler {
	return &eventServiceHandler{dbHandler: dbHandler}
}

func (es *eventServiceHandler) AddEvent(c echo.Context) error {

	var event model.Event
	if err := c.Bind(&event); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := es.dbHandler.AddEvent(event); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Event Added")
}

func (es *eventServiceHandler) FindEvent(c echo.Context) error {

	searchCriteria := c.Param("SearchCriteria")
	search := c.Param("search")
	var event model.Event
	var err error
	switch strings.ToLower(searchCriteria) {
	case "name":
		event, err = es.dbHandler.FindEventByName(search)
	case "id":
		event, err = es.dbHandler.FindEventByID(search)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError,
			fmt.Sprintf("error %s  search by  %s with value %s ", err.Error(), searchCriteria, search))
	}

	return c.JSON(http.StatusOK, event)

	return c.String(http.StatusOK, "Find Event")
}

func (es *eventServiceHandler) FindAll(c echo.Context) error {

	events, err := es.dbHandler.FindAllEvents()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, events)

}
