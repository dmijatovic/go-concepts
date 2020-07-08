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

//PostUpload uploads file using POST method with binary format
func (f *File) PostUpload(rw http.ResponseWriter, r *http.Request) {
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

//MultipartUpload handles multipart form uploads
func (f *File) MultipartUpload(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Println("Invalid multipart form")
		http.Error(rw, "Invalid multipart form", http.StatusBadRequest)
	}
	// extract FromFile
	fl, mh, e := req.FormFile("file")
	if e != nil {
		http.Error(rw, "Failed to retreive file from form", http.StatusBadRequest)
	}
	//extract file name
	fn := mh.Filename
	if fn != "" {
		f.log.Println("Uploading...", fn)
		f.store.Save(fn, fl)
	} else {
		http.Error(rw, "Failed to retreive Filename from form", http.StatusBadRequest)
	}
}
