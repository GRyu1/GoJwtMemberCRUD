package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		assert.Nil(t, err, "요청 만들기 오류")
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message":"pong"}`, rr.Body.String())
}
