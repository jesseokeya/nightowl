package services

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Folder is a json representation ot the resources (videos/images) found in a folder
type Folder struct {
	Path string // absolute path to the folder
}

// File is the description of the file found in the folder
type File struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Size      int64     `json:"size"`
	ParentDir string    `json:"parentDirectory"`
	SelfLink  string    `json:"selfLink"`
	Created   time.Time `json:"created"`
}

// Interprete gives us insights with a json representation of how a particular folder is structured
func (f *Folder) Interprete() []File {
	var files []File

	root := f.Path

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		file := f.extractFiles(info, path, root)
		if (File{}) != file {
			files = append(files, file)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

// DeleteFile deletes a file from the folder directory
func (f *Folder) DeleteFile(fileName string) ([]File, error) {
	var files []File
	root := f.Path

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() == fileName {
			os.Remove(path)
			log.Printf("file %s was successfully deleted", fileName)
		} else {
			file := f.extractFiles(info, path, root)
			if (File{}) != file {
				files = append(files, file)
			}
		}
		return nil
	})

	return files, err
}

// extractFiles isa aprivate helper method to get and build the file json object
func (f *Folder) extractFiles(info os.FileInfo, path string, root string) File {
	var file File

	if !info.IsDir() {
		split := strings.Split(path, "/")
		parentDir := strings.Join(split[:len(split)-1], "/")
		link := strings.Replace(path, root, "", 1)
		file = File{
			ParentDir: parentDir,
			Name:      info.Name(),
			Size:      info.Size(),
			Path:      path,
			SelfLink:  "/" + link,
			Created:   info.ModTime(),
		}
	}

	return file
}
