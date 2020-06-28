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
		io.WriteString(w, page404())
	}
}

func homePage() string {
	return "This is home page"
}

func pageID(id string) string {
	return fmt.Sprintf("This is page %v", id)
}

func page404() string {
	return "404 - Page not found"
}
