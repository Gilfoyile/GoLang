package main

import (
	"io"
	"os"
	"net/http"
)
func main(){
	url := "https://golangcode.com/images/avatar.jpg"

	if err := DownloadFile("avatar.jpg" ,url); err != nil{
		panic(err)
	}
}

func DownloadFile(filepath string , url string) error {
	resp,err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// creating the file

	out,err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	//Write the body to file 
	_,err = io.Copy(out,resp.Body)
	return err
}
