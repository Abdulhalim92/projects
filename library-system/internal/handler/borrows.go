package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
)

func (h *MyHandler) GetBorrows(w http.ResponseWriter, r *http.Request) {
	borrows, err := h.service.ListBorrows()
	if err != nil {
		log.Printf("Failed to list borrows - service.ListBooks error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(borrows, "   ", "")
	if err != nil {
		log.Printf("Failed to encode borrows - json MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write response - w.Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetFilteredBorrows(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	filterEncoded, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body - io.ReadAll error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	var filter *model.BorrowFilter
	err = json.Unmarshal(filterEncoded, &filter)
	if err != nil {
		log.Printf("Failed to decode request body - json.Unmarshal error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	borrows, err := h.service.ListFilteredBorrows(filter)
	if err != nil {
		log.Printf("Failed to get filtered borrows - service.GetFilteredBorrows error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(borrows, "   ", "")
	if err != nil {
		log.Printf("Failed to encode data - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write response - w.Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) AddBorrow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	borrowEncoded, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Add Borrow - io.ReadAll error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var borrow *model.Borrow
	err = json.Unmarshal(borrowEncoded, &borrow)
	if err != nil {
		log.Printf("Failed to Add Borrow - json.Unmarshal error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	borrow, err = h.service.CreateBorrow(borrow)
	if err != nil {
		log.Printf("Failed to Add Borrow - service.CreateBorrow error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(borrow, "   ", "")
	if err != nil {
		log.Printf("Failed to Add Borrow - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Add Borrow - Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) ReturnBorrow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	borrowEncoded, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Add Borrow - io.ReadAll error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var borrow *model.Borrow
	err = json.Unmarshal(borrowEncoded, &borrow)
	if err != nil {
		log.Printf("Failed to Add Borrow - json.Unmarshal error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	borrow, err = h.service.ReturnBorrow(borrow)
	if err != nil {
		log.Printf("Failed to Add Borrow - service.CreateBorrow error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(borrow, "   ", "")
	if err != nil {
		log.Printf("Failed to Add Borrow - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Add Borrow - Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
