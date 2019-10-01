package netutil

import (
	"io"
	"net/http"
	"os"
)

func Download(url, file string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	file1, err := os.Create(file)
	if err != nil {
		return err
	}
	defer file1.Close()

	_, err = io.Copy(file1, res.Body)
	if err != nil {
		return err
	}

	return nil
}
