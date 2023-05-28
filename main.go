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

	baseImgPath := "assets/nellfs/nellfs.png"
	baseImg, err := parseImage(baseImgPath)
	if err != nil {
		http.Error(w, "Failed to load base image", http.StatusInternalServerError)
		return
	}

	var overlayOptions OverlayOptions
	err = json.NewDecoder(r.Body).Decode(&overlayOptions)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	overlayOptionsPathMap := map[string]string{
		"red_hat":     "assets/items/red_hat.png",
		"glasses":     "assets/items/glasses.png",
		"anime_eyes":  "assets/items/anime_eyes.png",
		"tie":         "assets/items/tie.png",
		"wizard_hat":  "assets/items/wizard_hat.png",
		"bow_tie":     "assets/items/bow_tie.png",
		"halo":        "assets/items/halo.png",
		"mustache":    "assets/items/mustache.png",
		"red_glasses": "assets/items/red_glasses.png",
		"cat_ears":    "assets/items/cat_ears.png",
	}

	for key, path := range overlayOptionsPathMap {
		data := getOverlayOption(overlayOptions, key)
		if data != nil {
			overlayImg, err := parseImage(path)
			if err != nil {
				log.Printf("Failed to load overlay image for key '%s': %v", key, err)
				continue
			}
			overlay := imaging.Resize(overlayImg, data.W, data.H, imaging.NearestNeighbor)
			baseImg = imaging.Overlay(baseImg, overlay, image.Pt(data.X, data.Y), 1.0)
		}
	}

	resultPath := "result.png"
	err = imaging.Save(baseImg, resultPath)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
	}

	fmt.Println("Image overlay complete. Result saved as result.png.")
}

func getOverlayOption(options OverlayOptions, key string) *ImageData {
	switch key {
	case "red_hat":
		return options.RedHat
	case "glasses":
		return options.Glasses
	case "anime_eyes":
		return options.AnimeEyes
	case "tie":
		return options.Tie
	case "wizard_hat":
		return options.WizardHat
	case "bow_tie":
		return options.BowTie
	case "halo":
		return options.Halo
	case "mustache":
		return options.Mustache
	case "red_glasses":
		return options.RedGlasses
	case "cat_ears":
		return options.CatEars
	}
	return nil
}
