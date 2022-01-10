package test

import (
	"net/http"
	"testing"

	"github.com/Lukpier/mantil-sample-app/api/weather"
	"github.com/gavv/httpexpect"
)

func TestPing(t *testing.T) {
	api := httpexpect.New(t, apiURL)

	api.POST("/weather").
		Expect().
		Status(http.StatusOK).
		Body().Equal("Hello!")

	req := weather.Request{StationId: "10637"}
	api.POST("/weather/get").
		WithJSON(req).
		Expect().
		ContentType("application/json").
		Status(http.StatusOK)

	// method which don't exists
	api.POST("/weather/fake").
		Expect().
		Status(http.StatusNotImplemented)
}
