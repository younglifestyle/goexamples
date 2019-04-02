package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// function body of a http.HandlerFunc
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+1024)
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// parse text field
	text := make([]byte, 512)
	p, err := reader.NextPart()
	// one more field to parse, EOF is considered as failure here
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if p.FormName() != "text_field" {
		http.Error(w, "text_field is expected", http.StatusBadRequest)
		return
	}
	_, err = p.Read(text)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// parse file field
	p, err = reader.NextPart()
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if p.FormName() != "file_field" {
		http.Error(w, "file_field is expected", http.StatusBadRequest)
		return
	}
	buf := bufio.NewReader(p)
	sniff, _ := buf.Peek(512)
	contentType := http.DetectContentType(sniff)
	if contentType != "application/zip" {
		http.Error(w, "file type not allowed", http.StatusBadRequest)
		return
	}

	f, err := ioutil.TempFile("", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	var maxSize int64 = 32 << 20
	lmt := io.MultiReader(buf, io.LimitReader(p, maxSize-511))
	written, err := io.Copy(f, lmt)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if written > maxSize {
		os.Remove(f.Name())
		http.Error(w, "file size over limit", http.StatusBadRequest)
		return
	}
	// schedule for other stuffs (s3, scanning, etc.)
}

func main() {

}
