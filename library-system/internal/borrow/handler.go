package borrow

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
	"time"
)

type BorrowHandler struct {
	mux           *http.ServeMux
	borrowService *Service
}

func NewBorrowHandler(mux *http.ServeMux, s *Service) *BorrowHandler {
	return &BorrowHandler{
		mux:           mux,
		borrowService: s,
	}
}

func (h *BorrowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for borrows to the given ServeMux.
//
// It adds the following routes:
//
// - GET /borrows: GetBorrows
// - GET /borrows/:id: GetBorrowByID
// - POST /borrows/add: AddBorrow
// - DELETE /borrows/delete/:id: DeleteBorrow
func (h *BorrowHandler) InitRoutes() {
	h.mux.HandleFunc("/borrows", h.GetBorrows)
	h.mux.HandleFunc("/borrows/:id", h.GetBorrowByID)
	h.mux.HandleFunc("/borrows/add", h.AddBorrow)
	h.mux.HandleFunc("/borrows/delete/:id", h.DeleteBorrow)
}

func (h *BorrowHandler) GetBorrows(w http.ResponseWriter, r *http.Request) {
	borrows, err := h.borrowService.ListBorrows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(borrows, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling borrows: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *BorrowHandler) GetBorrowByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	borrowID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	borrow, err := h.borrowService.FindBorrow(borrowID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonBorrow, err := json.MarshalIndent(borrow, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling borrow: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBorrow)
} // TODO: implement

func (h *BorrowHandler) AddBorrow(w http.ResponseWriter, r *http.Request) {
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

	borrowDate, err := time.Parse("yy/mm/dd", r.FormValue("borrow_date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	returnDate, err := time.Parse("yy/mm/dd", r.FormValue("return_date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var borrow = &model.Borrow{
		UserId:     userID,
		BookId:     bookID,
		BorrowDate: borrowDate,
		ReturnDate: returnDate,
	}

	borrow, err = h.borrowService.br.AddBorrow(borrow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBorrow, err := json.MarshalIndent(borrow, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBorrow)

} // TODO: implement

func (h *BorrowHandler) DeleteBorrow(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	borrowID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.borrowService.RemoveBorrow(borrowID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Borrow with id:%d removed successfully", borrowID)
	w.Write([]byte(response))
} // TODO: implement
