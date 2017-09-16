package api

import "net/http"

// Flower type
type Flower struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (fl *Flower) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing
	return nil
}
