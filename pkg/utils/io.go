package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

var imageFileExtensionContentTypeMap = map[string]string{
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"gif":  "image/gif",
	"webp": "image/webp",
}


func GetImageContentType(fileExtension string) string {
	contentType, exists := imageFileExtensionContentTypeMap[fileExtension]

	if !exists {
		return ""
	}

	return contentType
}


func ListFileNamesWithPrefixAndSuffix(path string, prefix string, suffix string) []string {
	dir, err := os.Open(path)

	if err != nil {
		return nil
	}

	fileInfos, err := dir.Readdir(0)

	if err != nil {
		return nil
	}

	var fileNames []string

	for i := 0; i < len(fileInfos); i++ {
		fileInfo := fileInfos[i]

		if !fileInfo.IsDir() &&
			strings.HasPrefix(fileInfo.Name(), prefix) &&
			strings.HasSuffix(fileInfo.Name(), suffix) {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	return fileNames
}


func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}


func WriteFile(path string, data []byte) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	n, err := file.Write(data)

	if err == nil && n < len(data) {
		return io.ErrShortWrite
	}

	return err
}


func GetFileNameWithoutExtension(path string) string {
	if path == "" {
		return ""
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			path = path[i+1:]
			break
		}
	}

	if path == "" {
		return ""
	}

	extension := filepath.Ext(path)

	if len(extension) < 1 {
		return path
	}

	return path[0 : len(path)-len(extension)]
}


func GetFileNameExtension(path string) string {
	extension := filepath.Ext(path)

	if len(extension) < 1 || extension[0] != '.' {
		return extension
	}

	return extension[1:]
}
