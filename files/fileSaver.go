package files

import (
	"errors"
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
	_, err = file.Write(fileData)
	if err != nil {
		return err
	}
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
