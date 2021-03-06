package main

import (
	"PBLOG/router"
	"fmt"
	"net/http"
	"time"
)

func main() {
	r := router.RouterInit()
	r.Static("./static", "./static")
	r.Static("./edit", "./static/editormd")
	r.LoadHTMLFiles("./static/blog.html", "./static/index.html", "./static/add.html", "./static/edit.html")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8000),
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
