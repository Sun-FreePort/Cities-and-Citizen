package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSquareInfoSuccess(t *testing.T) {
	tearDown()
	setup()
	//username := "user2"
	url := "/api/articles"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-type", "application/json")

	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := io.ReadAll(resp.Body)
		var aa SquareInfoResp
		err := json.Unmarshal(body, &aa)
		assert.NoError(t, err)

		assert.Equal(t, int16(2), aa.SpeechCount)
	}
}
