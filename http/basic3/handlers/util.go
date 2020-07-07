package handlers

import (
	"log"
	"net/http"
	"regexp"
)

// ExtractIDFromURL wil extract id
func ExtractIDFromURL(r *http.Request) string {
	rx := regexp.MustCompile(`/([0-9 a-z])`)
	g := rx.FindAllStringSubmatch(r.URL.Path, -1)

	log.Println("Id extrated:", g)
	return g[0][1]
}
