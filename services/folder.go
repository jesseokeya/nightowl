package services

import (
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
		if !info.IsDir() {
			split := strings.Split(path, "/")
			parentDir := strings.Join(split[:len(split)-1], "/")
			link := strings.Replace(path, root, "", 1)
			file := File{
				ParentDir: parentDir,
				Name:      info.Name(),
				Size:      info.Size(),
				Path:      path,
				SelfLink:  os.Getenv("STORE") + link,
				Created:   info.ModTime(),
			}
			files = append(files, file)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
