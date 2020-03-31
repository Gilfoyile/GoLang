package main
import (
	"net/http"
)

func main(){
	http.HandleFunc("/download",pic)
	http.ListenAndServe(":8080",nil)
}
func pic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w,req,"avatar.jpg")
}

