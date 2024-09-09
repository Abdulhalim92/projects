package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
	"strings"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.service.ListReviews()
	if err != nil {
		log.Printf("ListReviews - h.service.ListReviews error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviews - reviews: %v", reviews)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(reviews, "", "    ")
	if err != nil {
		log.Printf("GetReviews - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviews - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) AddReview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateReview - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateReview - incoming request: %v", string(data))

	var review model.Review
	err = json.Unmarshal(data, &review)
	if err != nil {
		log.Printf("CreateReview - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateReview - data after unmarshalling: %v", review)

	createReview, err := h.service.CreateReview(&review)
	if err != nil {
		log.Printf("CreateReview - h.service.CreateReview error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateReview - created review: %v", createReview)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(createReview, "", "    ")
	if err != nil {
		log.Printf("CreateReview - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateReview - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reviews/get/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReview - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	review, err := h.service.GetReviewById(id)
	if err != nil {
		log.Printf("GetReview - h.service.GetReviewByID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReview - review: %v", review)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(review, "", "    ")
	if err != nil {
		log.Printf("GetReview - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReview - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetReviewsByBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reviews/get/book/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReviewsByBookID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reviews, err := h.service.ListReviewsByBookId(id)
	if err != nil {
		log.Printf("GetReviewsByBookID - h.service.GetReviewsByBookID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByBookID - reviews: %v", reviews)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(reviews, "", "    ")
	if err != nil {
		log.Printf("GetReviewsByBookID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByBookID - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetReviewsByUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reviews/get/user/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReviewsByUserID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reviews, err := h.service.ListReviewsByUserId(id)
	if err != nil {
		log.Printf("GetReviewsByUserID - h.service.GetReviewsByUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByUserID - reviews: %v", reviews)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(reviews, "", "    ")
	if err != nil {
		log.Printf("GetReviewsByUserID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByUserID - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetReviewsByFilter(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var filter model.ReviewFilter

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("GetReviewsByFilter - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByFilter - incoming request: %v", string(data))

	err = json.Unmarshal(data, &filter)
	if err != nil {
		log.Printf("GetReviewsByFilter - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByFilter - data after unmarshalling: %v", filter)

	reviews, err := h.service.GetReviewsByFilter(filter)
	if err != nil {
		log.Printf("GetReviewsByFilter - h.service.GetReviewsByFilter error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByFilter - reviews: %v", reviews)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(reviews, "", "    ")
	if err != nil {
		log.Printf("GetReviewsByFilter - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetReviewsByFilter - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("EditReview - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditReview - incoming request: %v", string(data))

	var review model.Review
	err = json.Unmarshal(data, &review)
	if err != nil {
		log.Printf("EditReview - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditReview - data after unmarshalling: %v", review)

	updatedReview, err := h.service.EditReview(&review)
	if err != nil {
		log.Printf("EditReview - h.service.EditReview error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditReview - updated review: %v", updatedReview)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(updatedReview, "", "    ")
	if err != nil {
		log.Printf("EditReview - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditReview - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reviews/delete/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteReview - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.service.RemoveReviewById(id)
	if err != nil {
		log.Printf("DeleteReview - h.service.DeleteReview error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted review"))
}
