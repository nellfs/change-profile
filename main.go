package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	http.HandleFunc("/", handleOverlay)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleOverlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	baseImg, _ := parseImage("assets/nellfs/nellfs.png")

	decoder := json.NewDecoder(r.Body)
	var request OverlayOptions
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	img, err := parseImage("assets/items/anime_eyes.png")
	if err != nil {
		log.Fatalf("Unable to open image  %v", err)
	}

	baseImg = createImage(baseImg, img, request.AnimeEyes)
	baseImg = createImage(baseImg, img, request.Tie)

	err = imaging.Save(baseImg, "./result.png")
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	fmt.Println("Image overlay complete. Result saved as result.png.")
}

func createImage(baseImg image.Image, overlayImg image.Image, data ImageData) image.Image {
	overlay := imaging.Resize(overlayImg, data.W, data.H, imaging.NearestNeighbor)
	result := imaging.Overlay(baseImg, overlay, image.Pt(data.X, data.Y), 1.0)

	return result
}
