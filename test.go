package main

import (
	"fmt"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	url := "http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=Wiki.png")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	//stream the body to the client without fully loading it into memory
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
