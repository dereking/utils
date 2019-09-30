package netutil

import(
 
    "io"
    "net/http"
    "os"
)

func Download(url, file string) error {
	res, _ := http.Get(url)
    file1, _ := os.Create(file)
    io.Copy(file1, res.Body) 
	
	return nil
}