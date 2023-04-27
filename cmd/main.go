package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
)

const chunkSize = 1024 // 1 kilobyte = 1024 bytes

func main() {
	url := "http://example.com/archive.tar.gz" // replace with the actual archive URL

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if h.Typeflag != tar.TypeReg {
			// skip non-regular files
			continue
		}

		c := make(chan []byte)

		// concurrent downloader
		go func(src io.Reader) {
			buf := make([]byte, chunkSize)
			for {
				n, err := src.Read(buf)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Println(err)
					break
				}
				c <- buf[:n]
			}
			close(c)
		}(tr)

		// acceptor
		for chunk := range c {
			// process the chunk
			fmt.Println(string(chunk))
		}
	}
}
