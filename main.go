// Application to fetch and store images from Grafana's panel renderer
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// run the server
func main() {
	var (
		port      = flag.Int("port", 8080, "grafana-images listening port")
		imageHost = flag.String("imageHost", "http://grafana.example.com/saved-images", "host for the saved images")
		imagePath = flag.String("imagePath", "/opt/saved-images", "location on disk where images will be saved")
	)
	flag.Parse()

	http.HandleFunc("/grafana-images", GrafanaImagesHandler(*imageHost, *imagePath))

	log.Println("grafana-images: listening on port", *port, "; image host", *imageHost, "; image path", *imagePath)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
