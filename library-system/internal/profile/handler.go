package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
)

type ProfileHandler struct {
	mux            *http.ServeMux
	profileService *Service
}

func NewProfileHandler(mux *http.ServeMux, s *Service) *ProfileHandler {
	return &ProfileHandler{
		mux:            mux,
		profileService: s,
	}
}

func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for profiles to the given ServeMux.
//
// It adds the following routes:
//
// - GET /profiles: GetProfiles
// - GET /profiles/:id: GetProfileByID
// - POST /profiles/add: AddProfile
// - DELETE /profiles/delete/:id: DeleteProfile
func (h *ProfileHandler) InitRoutes() {
	h.mux.HandleFunc("/profiles", h.GetProfiles)
	h.mux.HandleFunc("/profiles/:id", h.GetProfileByID)
	h.mux.HandleFunc("/profiles/add", h.AddProfile)
	h.mux.HandleFunc("/profiles/delete/:id", h.DeleteProfile)
}

func (h *ProfileHandler) GetProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.profileService.ListProfiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(profiles, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling profiles: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	profileID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	profile, err := h.profileService.FindProfile(profileID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonProfile, err := json.MarshalIndent(profile, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling profile: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProfile)
} // TODO: implement

func (h *ProfileHandler) AddProfile(w http.ResponseWriter, r *http.Request) {
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
	email := r.FormValue("email")
	addr := r.FormValue("addr")

	var profile = &model.Profile{
		UserId:  userID,
		Email:   email,
		Address: addr,
	}

	profile, err = h.profileService.ur.AddProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonProfile, err := json.MarshalIndent(profile, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProfile)

} // TODO: implement

func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	profileID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.profileService.RemoveProfile(profileID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Profile with id:%d removed successfully", profileID)
	w.Write([]byte(response))
} // TODO: implement
