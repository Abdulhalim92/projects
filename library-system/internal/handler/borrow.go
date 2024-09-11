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

func (h *Handler) GetBorrows(w http.ResponseWriter, r *http.Request) {
	borrows, err := h.service.GetBorrows()
	if err != nil {
		log.Printf("GetBorrows - h.service.GetBorrows error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrows - borrows: %v", borrows)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(borrows, "", "    ")
	if err != nil {
		log.Printf("GetBorrows - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrows - data: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetBorrowByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrows/")

	if idStr == "" {
		log.Printf("GetBorrowByID - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	borrow, err := h.service.GetBorrowByID(id)
	if err != nil {
		log.Printf("GetBorrowByID - h.service.GetBorrowByID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByID - borrow: %v", borrow)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(borrow, "", "    ")
	if err != nil {
		log.Printf("GetBorrowByID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByID - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetBorrowByUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrows/user/")
	if idStr == "" {
		log.Printf("GetBorrowByUser - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByUser - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	borrows, err := h.service.GetBorrowsByUser(id)
	if err != nil {
		log.Printf("GetBorrowByUser - h.service.GetBorrowsByUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByUser - borrows: %v", borrows)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(borrows, "", "    ")
	if err != nil {
		log.Printf("GetBorrowByUser - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByUser - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetBorrowByBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrows/book/")
	if idStr == "" {
		log.Printf("GetBorrowByBook - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByBook - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	borrows, err := h.service.GetBorrowsByBook(id)
	if err != nil {
		log.Printf("GetBorrowByBook - h.service.GetBorrowsByBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByBook - borrows: %v", borrows)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(borrows, "", "    ")
	if err != nil {
		log.Printf("GetBorrowByBook - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBorrowByBook - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) CreateBorrow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateBorrow - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateBorrow - incoming request: %v\n", string(body))

	var borrow model.Borrow

	err = json.Unmarshal(body, &borrow)
	if err != nil {
		log.Printf("CreateBorrow - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateBorrow - data after unmarshalling: %v", borrow)

	createdBorrow, err := h.service.CreateBorrow(&borrow)
	if err != nil {
		log.Printf("CreateBorrow - h.service.CreateBorrow error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateBorrow - createdBorrow: %v", createdBorrow)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(createdBorrow, "", "    ")
	if err != nil {
		log.Printf("CreateBorrow - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateBorrow - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrows/return/")
	if idStr == "" {
		log.Printf("ReturnBook - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ReturnBook - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.ReturnBook(id)
	if err != nil {
		log.Printf("ReturnBook - h.service.ReturnBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book returned"))
}
