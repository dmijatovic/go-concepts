package handlers

import (
	"log"
	"net/http"
)

//Hello creates handler
// which needs to implement interface
type Hello struct {
	l *log.Logger
}

// NewHello return pointer to a new instance of Hello handler
// with desired logger.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")

	rw.Write([]byte("Basic hello world"))

}
