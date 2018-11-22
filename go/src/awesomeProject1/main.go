package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/blank", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"./New Site/blank.html")
	})
	fs:=http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/theme.css",fs))
	http.ListenAndServe(":8181",nil)
}