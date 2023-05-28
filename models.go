package main

type ImageData struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type OverlayOptions struct {
	RedHat     *ImageData `json:"red_hat"`
	Glasses    *ImageData `json:"glasses"`
	AnimeEyes  *ImageData `json:"anime_eyes"`
	Tie        *ImageData `json:"tie"`
	WizardHat  *ImageData `json:"wizard_hat"`
	BowTie     *ImageData `json:"bow_tie"`
	Halo       *ImageData `json:"halo"`
	Mustache   *ImageData `json:"mustache"`
	RedGlasses *ImageData `json:"red_glasses"`
}
