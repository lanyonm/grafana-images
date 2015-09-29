package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GrafanaImage struct {
	Url string `json:"imageUrl"`
}

func GrafanaImagesHandler(imageHost string, imagePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var image GrafanaImage

		if req.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			w.Header().Set("Allow", "POST")
			return
		}

		if err := json.NewDecoder(req.Body).Decode(&image); err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		token := req.Header.Get("Authorization")
		log.Println("received request:", image.Url, "token:", token)

		// Fetch image
		var client = http.Client{}
		req, err := http.NewRequest("GET", image.Url, nil)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", token)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("http.Get -> %v", err)
			w.WriteHeader(500)
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("ioutil.ReadAll -> %v", err)
			w.WriteHeader(500)
			return
		}

		// Save image
		fileName := fmt.Sprintf("%x.png", md5.Sum(data))
		resp.Body.Close()
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s", imagePath, fileName), data, 0666)
		if err != nil {
			log.Fatalf("ioutil.WriteFile -> %v", err)
			w.WriteHeader(500)
			return
		}

		// Return image location
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"pubImg\":\"%s/%s\"}", imageHost, fileName)
	}
}
