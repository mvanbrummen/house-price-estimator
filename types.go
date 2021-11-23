package main

import "github.com/leekchan/accounting"

type AccessResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

type SuggestResponse struct {
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

type Suggestion struct {
	PropertyId int    `json:"propertyId,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

type ValuationResponse struct {
	Confidence   string `json:"confidence,omitempty"`
	Estimate     int    `json:"estimate,omitempty"`
	HighEstimate int    `json:"highEstimate,omitempty"`
	LowEstimate  int    `json:"lowEstimate,omitempty"`
}

type ImageryResponse struct {
	DefaultImage       Image   `json:"defaultImage,omitempty"`
	FloorPlanImageList []Image `json:"floorPlanImageList,omitempty"`
	SecondaryImageList []Image `json:"secondaryImageList,omitempty"`
}

type Image struct {
	ThumbnailPhotoUrl string `json:"thumbnailPhotoUrl,omitempty"`
	MediumPhotoUrl    string `json:"mediumPhotoUrl,omitempty"`
	LargePhotoUrl     string `json:"largePhotoUrl,omitempty"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	Msg string `json:"msg,omitempty"`
}

type Valuation struct {
	LowEstimate        string
	Estimate           string
	HighEstimate       string
	Confidence         string
	Address            string
	DefaultImageUrl    string
	SecondaryImageUrls []string
}

func mapValuation(v *ValuationResponse, address string) *Valuation {
	ac := accounting.Accounting{Symbol: "$", Precision: 0}
	return &Valuation{
		LowEstimate:  ac.FormatMoney(v.LowEstimate),
		Estimate:     ac.FormatMoney(v.Estimate),
		HighEstimate: ac.FormatMoney(v.HighEstimate),
		Confidence:   v.Confidence,
		Address:      address,
	}
}
