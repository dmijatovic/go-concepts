package files

import "io"

// Storage defines interface for saving files
type Storage interface {
	Save(path string, file io.Reader) error
}
