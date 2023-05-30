package main

type Crop struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Size int `json:"size"`
}

type OverlayData struct {
	X int     `json:"x"`
	Y int     `json:"y"`
	W int     `json:"w"`
	H int     `json:"h"`
	R float64 `json:"r"`
}

type OverlayOptions struct {
	RedHat     *OverlayData `json:"red_hat"`
	Glasses    *OverlayData `json:"glasses"`
	AnimeEyes  *OverlayData `json:"anime_eyes"`
	Tie        *OverlayData `json:"tie"`
	WizardHat  *OverlayData `json:"wizard_hat"`
	BowTie     *OverlayData `json:"bow_tie"`
	Halo       *OverlayData `json:"halo"`
	Mustache   *OverlayData `json:"mustache"`
	RedGlasses *OverlayData `json:"red_glasses"`
	CatEars    *OverlayData `json:"cat_ears"`
}

type ImageData struct {
	Image    int            `json:"image"`
	Rotation float64        `json:"rotation"`
	Crop     *Crop          `json:"crop"`
	Overlay  OverlayOptions `json:"overlay"`
}
