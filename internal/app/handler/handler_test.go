package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnHelloWorldWhentheHandlerIsInvoked(t *testing.T) {

	engine := gin.Default()

	engine.GET("/test/hello", SayHello)

	request, err := http.NewRequest(http.MethodGet, "/test/hello", nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	bytes, err := io.ReadAll(responseRecoder.Body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecoder.Result().StatusCode)
	assert.Equal(t, "hello world!", string(bytes))

}
