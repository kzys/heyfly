package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed page.html
var pageHTML string

type PageData struct {
	Request *http.Request
	Environ []string
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("request = %+v", req)
	t, err := template.New("page").Parse(pageHTML)
	if err != nil {
		log.Printf("failed to parse: %s", err)
		return
	}

	w.Header().Set("content-type", "text/html; charset=utf-8")

	f, err := os.Open("/env.txt")
	if err != nil {
		log.Printf("failed to open: %s", err)
		return
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		log.Printf("failed to open: %s", err)
		return
	}

	data := PageData{
		Request: req,
		Environ: os.Environ(),
	}
	err = t.Execute(w, &data)
	if err != nil {
		log.Printf("failed to execute: %s", err)
	}
}

func main() {
	port := 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("listen %d", port)
	log.Fatal(s.ListenAndServe())
}
