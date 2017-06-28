package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"net/http"
)

var (
	server *httptest.Server
)

type Response struct {
	Content string
	Code    int
}

// HTTP Test Init
func Test_Init(t *testing.T) {
	logfile, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	assert.Nil(t, err, "")
	h := Handler{}
	h.Init()
	server = httptest.NewServer(handlers.CombinedLoggingHandler(logfile, http.DefaultServeMux))
	testUrl = server.URL
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	var a string = "Heelooo?"
	var b string = "Heelooo?"

	assert.Equal(a, b, "the two words should be the smae.")

}