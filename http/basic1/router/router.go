package router

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Router can be any type that implements
// ServeHTTP method
type Router struct{}

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log request
	log.Printf("Route...%v", r.URL.Path)
	// route
	switch r.URL.Path {
	case "/page1", "/page2":
		io.WriteString(w, pageID(r.URL.Path[5:]))
	case "/":
		io.WriteString(w, homePage())
	default:
		page404(w)
	}
}

func homePage() string {
	return "This is home page"
}

func pageID(id string) string {
	return fmt.Sprintf("This is page %v", id)
}

func page404(rw http.ResponseWriter) {
	http.Error(rw, "404-Page not found", http.StatusNotFound)
}
