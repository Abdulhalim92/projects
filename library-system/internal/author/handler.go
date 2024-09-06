package author

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	mux           *http.ServeMux
	authorService *Service
}

func NewAuthorHandler(mux *http.ServeMux, s *Service) *AuthorHandler {
	return &AuthorHandler{
		mux:           mux,
		authorService: s,
	}
}

func (h *AuthorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for authors to the given ServeMux.
//
// It adds the following routes:
//
// - GET /authors: GetAuthors
// - GET /authors/:id: GetAuthorByID
// - POST /authors/add: AddAuthor
// - DELETE /authors/delete/:id: DeleteAuthor
func (h *AuthorHandler) InitRoutes() {
	h.mux.HandleFunc("/authors", h.GetAuthors)
	h.mux.HandleFunc("/authors/:id", h.GetAuthorByID)
	h.mux.HandleFunc("/authors/add", h.AddAuthor)
	h.mux.HandleFunc("/authors/delete/:id", h.DeleteAuthor)
}

func (h *AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.authorService.ListAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(authors, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling authors: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	authorID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	author, err := h.authorService.FindAuthor(authorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonAuthor, err := json.MarshalIndent(author, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling author: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonAuthor)
} // TODO: implement

func (h *AuthorHandler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	name := r.FormValue("name")
	bio := r.FormValue("bio")

	var author = &model.Author{
		Name:      name,
		Biography: bio,
	}

	author, err = h.authorService.ar.AddAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonAuthor, err := json.MarshalIndent(author, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAuthor)

} // TODO: implement

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	authorID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.authorService.RemoveAuthor(authorID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Author with id:%d removed successfully", authorID)
	w.Write([]byte(response))
} // TODO: implement
