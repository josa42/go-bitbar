package bitbar

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
)

type Image struct {
	path   string
	width  int
	height int
}

func newImage(path string) *Image {
	return &Image{path: path}
}

func (i *Image) string(key string) string {

	path := i.path
	if path != "" {

		if strings.HasPrefix(path, "http") {
			path = downloadImage(path)
		}

		data, _ := ioutil.ReadFile(path)
		return fmt.Sprintf("%s=%s", key, base64.StdEncoding.EncodeToString(data))
	} else {

	}

	return ""
}

func downloadImage(url string) string {

	usr, _ := user.Current()

	dir := path.Join(usr.HomeDir, "Library", "Caches", "go-bitbar", "image")
	os.MkdirAll(dir, os.ModePerm)

	hashb := md5.Sum([]byte(url))
	hash := string(hex.EncodeToString(hashb[:]))

	file := path.Join(dir, hash)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		output, err := os.Create(file)
		if err != nil {
			return ""
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			return ""
		}
		defer response.Body.Close()

		_, errCopy := io.Copy(output, response.Body)
		if errCopy != nil {
			return ""
		}
	}

	return file
}
