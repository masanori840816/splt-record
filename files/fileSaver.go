package files

import (
	"errors"
	"log"
	"os"
	"path"
)

func SaveFile(saveDir string, fileName string, fileData []byte) error {
	err := createDirectory(saveDir)
	if err != nil {
		return err
	}
	filePath := path.Join(saveDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	saveLen, err := file.Write(fileData)
	if err != nil {
		return err
	}
	log.Println(saveLen)
	return nil
}
func createDirectory(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
