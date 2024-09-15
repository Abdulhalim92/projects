package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
)

func (h *MyHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	revs, err := h.service.ListReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(revs, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetFilteredReviews(w http.ResponseWriter, r *http.Request) {
	//
	//borrows, err := h.service.ListFilteredBorrows(filter)
	//if err != nil {
	//	log.Printf("Failed to get filtered borrows - service.GetFilteredBorrows error %v\n", err)
	//	http.Error(w, "", http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//data, err := json.MarshalIndent(borrows, "   ", "")
	//if err != nil {
	//	log.Printf("Failed to encode data - json.MarshalIndent error %v\n", err)
	//	http.Error(w, "", http.StatusInternalServerError)
	//	return
	//}
	//_, err = w.Write(data)
	//if err != nil {
	//	log.Printf("Failed to write response - w.Write error %v\n", err)
	//	http.Error(w, "", http.StatusInternalServerError)
	//	return
	//}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Get Filtered Reviews - io.ReadAll error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filter *model.ReviewFilter
	err = json.Unmarshal(data, &filter)
	if err != nil {
		log.Printf("Failed to Get Filtered Reviews - json.Unmarshal error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	revs, err := h.service.ListFilteredReviews(filter)
	if err != nil {
		log.Printf("Failed to Get Filtered Reviews - service.ListFilteredReviews error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(revs, "   ", "")
	if err != nil {
		log.Printf("Failed to Get Filtered Reviews - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Get Filtered Reviews - w.Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) AddReview(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Add Review - io.ReadAll error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var review *model.Review
	err = json.Unmarshal(data, &review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rev, err := h.service.CreateReview(review)
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(rev, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
