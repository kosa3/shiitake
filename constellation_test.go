package main

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestFetchFortuneTelling(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", BaseUrl+"20200601.json", func(req *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(200, ShiitakeResponse{
			Aries: Aries{
				Payload{
					Analysis: "analysis",
					Advice:   "advice",
					PowerUp:  "power up",
					CoolDown: "cool down",
				},
			},
		})
		if err != nil {
			return httpmock.NewStringResponse(500, ""), nil
		}

		return res, nil
	})
	mockTime := time.Date(2020, 6, 1, 1, 2, 3, 123456000, time.UTC)
	expected, _ := fetchFortuneTelling(mockTime)

	assert.Equal(t, expected.Aries.Payload.Analysis, "analysis")
}
