package handlers

import (
	"log"
	"net/http"
)

//Goodbye handler which MUST implement ServeHTTP method
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye return pointer to a new instance of Hello handler
// with desired logger.
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Goodbye route")

	rw.Write([]byte("Goodbye little princes..."))

}
