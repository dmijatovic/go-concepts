package files

import (
	"io"
	"os"
	"path/filepath"
)

//Local file upload object
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocalUpload create new instance for local file upload
func NewLocalUpload(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

//Save implements uploading interface define by Storage interface
func (l *Local) Save(path string, content io.Reader) error {

	// construct full path
	fp := filepath.Join(l.basePath, path)
	// ensure directory exists
	dir := filepath.Dir(fp)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	//delete previous file if exists
	err = deleteIfExists(fp)
	if err != nil {
		return err
	}
	//create new file
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	//write file to disk
	//ensure max file size met
	_, err = io.Copy(f, content)
	if err != nil {
		return err
	}
	return nil
}

func deleteIfExists(filename string) error {
	// delete if file aleady exists
	_, err := os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}
