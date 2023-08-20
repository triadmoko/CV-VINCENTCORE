package helpers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetFileContentType(dir string) (string, error) {
	file, err := os.Open(dir)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 512)

	_, err = file.Read(buf)

	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buf)

	return contentType, nil
}

func FileNameWithoutExtSliceNotation(fileName string) (string, string) {
	filename := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	ext := filepath.Ext(fileName)
	filename = UUID()
	return filename, ext
}

func GetIDStorageByUrl(urlPath string) (id string, err error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return urlPath, err
	}
	splitUrl := strings.Split(u.Path, "/")
	fileID := splitUrl[len(splitUrl)-1 :]
	splitFilename := strings.Split(fileID[0], ".")

	return splitFilename[0], nil
}

func GetFileUrl(bucket, object string) string {
	return LoadEnv("DOMAIN_FILE") + "/" + bucket + "/" + object
}
