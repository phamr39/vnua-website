package models


type Video struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoUrl    string `json:"video_url"`
}
