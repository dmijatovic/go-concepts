package handlers

import (
	"log"
	"net/http"

	"../files"
	"github.com/gorilla/mux"
)

// File strucure
type File struct {
	log   *log.Logger
	store files.Storage
}

// NewFile returns handler for uploading files
func NewFile(l *log.Logger) *File {
	//relative path from project root
	st, err := files.NewLocalUpload("./images", 20000)
	if err != nil {
		log.Panic(err)
	}
	return &File{log: l, store: st}
}

func (f *File) logRequest(r *http.Request) {
	f.log.Println(r.URL.Path + "..." + r.Method)
}

func (f *File) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fn := vars["filename"]
	// log request
	f.logRequest(r)

	if fn == "" {
		http.Error(rw, "Invalid filename", http.StatusBadRequest)
		return
	}

	// log upload
	f.log.Println("Uploading...", fn)
	// upload file
	f.store.Save(fn, r.Body)
}

// func (f *File) SaveFile(filename string, rw http.ResponseWriter, req *http.Request) {
// 	files.Save()
// }
