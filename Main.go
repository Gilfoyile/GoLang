package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	p := NewsAggPage{Title : "The Amazing News" , News : "Some News"}
	t,_ :=template.ParseFiles("basictemplating.html")
	t.Execute(w,p)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1> Hello There Umang Singh</h1>")
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/agg/",newsAggHandler)
	http.ListenAndServe(":3000",nil)
}

