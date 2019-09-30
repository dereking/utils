package netutil

import(
 
    "io"
    "net/http"
    "os"
)

func Download(url, file string) error {
	res, _ := http.Get(url)
    file, _ := os.Create(file)
    io.Copy(file, res.Body) 
	
	return nil
}