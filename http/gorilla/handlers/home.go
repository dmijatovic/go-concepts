package handlers

import (
	"log"
	"net/http"
)

// Home creates handler
// which needs to implement interface
type Home struct {
	l *log.Logger
}

// NewHome return pointer to a new instance of Home handler
// with desired logger.
func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (h *Home) logRequest(r *http.Request) {
	h.l.Println(r.URL.Path + "..." + r.Method)
}

func (h *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logRequest(r)
	rw.Write([]byte("Basic Home world"))
}
