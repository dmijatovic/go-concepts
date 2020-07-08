package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

//GzipHandler gzips files
type GzipHandler struct{}

func (g *GzipHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//get encoding
	enc := req.Header.Get("Accept-Encoding")
	if strings.Contains(enc, "gzip") {
		rw.Write([]byte("Its working"))
	}
}

// GzipMiddleware gzips returns
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		enc := req.Header.Get("Accept-Encoding")
		if strings.Contains(enc, "gzip") {
			rw.Write([]byte("Its working"))
		}
	}
}

//GzipResponseWritter write gzipped data
type GzipResponseWritter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

//NewGzipResponseWriter create gzip response writer
func NewGzipResponseWriter(rw http.ResponseWriter) *GzipResponseWritter {
	gw := gzip.NewWriter(rw)
	return &GzipResponseWritter{rw: rw, gw: gw}
}

//Header returns original response header
func (gw *GzipResponseWritter) Header() http.Header {
	return gw.rw.Header()
}

//Write gzip data back
func (gw *GzipResponseWritter) Write(d []byte) (int, error) {
	return gw.gw.Write(d)
}

//WriteHeader passes header to response writer
func (gw *GzipResponseWritter) WriteHeader(status int) {
	gw.rw.WriteHeader(status)
}

//Flush will flush all data
func (gw *GzipResponseWritter) Flush() {
	gw.gw.Flush()
	gw.gw.Close()
}
