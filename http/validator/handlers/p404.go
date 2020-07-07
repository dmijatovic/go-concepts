package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// P404 creates handler
// which needs to implement interface
type P404 struct {
	l *log.Logger
}

// NewP404 return pointer to a new instance of P404 handler
// with desired logger.
func NewP404(l *log.Logger) *P404 {
	return &P404{l}
}

func (h *P404) logRequest(r *http.Request) *P404 {
	h.l.Println(fmt.Println(r.URL.Path + "..." + r.Method))
	return h
}

func (h *P404) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logRequest(r)
	rw.Write([]byte("Basic P404 world"))
}
