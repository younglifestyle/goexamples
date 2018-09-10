package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"net/http"

	"github.com/gin-gonic/gin/ginS"
	"github.com/juju/ratelimit"
)

func main() {
	// Source holding 1MB
	src := bytes.NewReader(make([]byte, 1024*1024))
	// Destination
	dst := &bytes.Buffer{}

	// Bucket adding 100KB every second, holding max 100KB
	bucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)

	start := time.Now()

	// Copy source to destination, but wrap our reader with rate limited one
	io.Copy(dst, ratelimit.Reader(src, bucket))

	fmt.Printf("Copied %d bytes in %s\n", dst.Len(), time.Since(start))
}

func test() {
	ginS.StaticFile()

	http.ServeFile()
}
