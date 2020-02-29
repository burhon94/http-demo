package pkg

import (
	"log"
	"os"
	"path/filepath"
)

func FilesDir(dir string) (files []string, err error) {
	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) (err2 error) {
			if err != nil {
				err2 = err
				return err2
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return files, nil
}
