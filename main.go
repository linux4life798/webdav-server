// Minimal usage of the golang.org/x/net webdav server.
//
// Originally written by Claude:
// https://claude.site/artifacts/4bfb0c72-e6ec-4917-8185-b3836a7a24a1
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	port := flag.Int("port", 8080, "Port to serve WebDAV on")
	dir := flag.String("dir", ".", "Directory to serve via WebDAV")
	flag.Parse()

	handler := &webdav.Handler{
		FileSystem: webdav.Dir(*dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WebDAV error: %s, %s", r.Method, err)
			} else {
				log.Printf("WebDAV %s: %s", r.Method, r.URL)
			}
		},
	}

	http.Handle("/", handler)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting WebDAV server on http://localhost%s", addr)
	log.Printf("Serving directory: %s", *dir)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
