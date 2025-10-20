package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello Nikko")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	response := httptest.NewRecorder()

	HelloHandler(response, request)

	// res := response.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	assert.Equal(t, "Hello Nikko", string(body))

}
