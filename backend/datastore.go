// datastore.go
package backend

// TODO:
//      global hash value to avoid cross interaction between objects, just use reset
//		disallow dirs to be opened with new func
//

import (
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type DataStore interface {
	AddBytes(tag string, data []byte) (err error)
	AddBytesFromReader(tag string, reader *io.Reader) (err error)
	RemoveBytes(tag string) (err error)
}

type File struct {
	os.File
	original_name string
}

type FileStore struct {
	dump_directory string
}

func (file *File) OriginalName() (orig_name string) {
	return file.original_name
}

// NewFileStore creates a new FileStore structure with the specified dump directory for files,
// returns an error when that directory is inaccessible
func NewFileStore(dump_dir string) (store *FileStore, err error) {
	_, err = os.Stat(dump_dir)
	if err == nil {
		return &FileStore{
			dump_directory: dump_dir,
		}, nil
	}
	return nil, err
}

func (store *FileStore) DumpDirectory() (dir string) {
	return store.dump_directory
}

func (store *FileStore) OpenFile(file_name string, flag int, perm os.FileMode) (file *File, err error) {
	full_path := filepath.Join(store.dump_directory, strconv.FormatUint(hashString(file_name), 10))
	os_file, err := os.OpenFile(full_path, flag, perm)
	return &File{
		File:          *os_file,
		original_name: file_name,
	}, err
}

func (store *FileStore) DeleteFile(file_name string) (err error) {
	full_path := filepath.Join(store.dump_directory, strconv.FormatUint(hashString(file_name), 10))
	return os.Remove(full_path)
}

func hashString(text string) (hash uint64) {
	hasher := fnv.New64()
	io.WriteString(hasher, text)
	return hasher.Sum64()
}

//func (store *FileStore) AddBytes(tag string, data []byte) (err error) {

//}

//func (store *FileStore) AddBytesFromReader(tag string, reader *io.Reader) (err error) {

//}

//func (store *FileStore) RemoveBytes(tag string) (err error) {

//}
