package review

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
	"time"
)

type ReviewHandler struct {
	mux           *http.ServeMux
	reviewService *Service
}

func NewReviewHandler(mux *http.ServeMux, s *Service) *ReviewHandler {
	return &ReviewHandler{
		mux:           mux,
		reviewService: s,
	}
}

func (h *ReviewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for reviews to the given ServeMux.
//
// It adds the following routes:
//
// - GET /reviews: GetReviews
// - GET /reviews/:id: GetReviewByID
// - POST /reviews/add: AddReview
// - DELETE /reviews/delete/:id: DeleteReview
func (h *ReviewHandler) InitRoutes() {
	h.mux.HandleFunc("/reviews", h.GetReviews)
	h.mux.HandleFunc("/reviews/:id", h.GetReviewByID)
	h.mux.HandleFunc("/reviews/add", h.AddReview)
	h.mux.HandleFunc("/reviews/delete/:id", h.DeleteReview)
}

func (h *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.reviewService.ListReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(reviews, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling reviews: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *ReviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reviewID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	review, err := h.reviewService.FindReview(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonReview, err := json.MarshalIndent(review, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling review: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonReview)
} // TODO: implement

func (h *ReviewHandler) AddReview(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	bookID, err := strconv.Atoi(r.FormValue("book_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reviewText := r.FormValue("review_text")
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reviewDate, err := time.Parse("yy/mm/dd", r.FormValue("review_date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var review = &model.Review{
		UserId:     userID,
		BookId:     bookID,
		ReviewText: reviewText,
		Rating:     float64(rating),
		ReviewDate: reviewDate,
	}

	review, err = h.reviewService.rr.AddReview(review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonReview, err := json.MarshalIndent(review, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReview)

} // TODO: implement

func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reviewID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.reviewService.RemoveReview(reviewID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Review with id:%d removed successfully", reviewID)
	w.Write([]byte(response))
} // TODO: implement
