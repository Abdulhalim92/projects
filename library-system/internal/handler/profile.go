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

func (h *Handler) ListProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.service.ListProfiles()
	if err != nil {
		log.Printf("ListProfiles - h.service.ListProfiles error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("ListProfiles - profiles: %v", profiles)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(profiles, "", "    ")
	if err != nil {
		log.Printf("ListProfiles - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("ListProfiles - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateProfile - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateProfile - incoming request: %v", string(data))

	var profile model.Profile
	err = json.Unmarshal(data, &profile)
	if err != nil {
		log.Printf("CreateProfile - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateProfile - data after unmarshalling: %v", profile)

	createProfile, err := h.service.CreateProfile(&profile)
	if err != nil {
		log.Printf("CreateProfile - h.service.CreateProfile error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateProfile - created profile: %v", createProfile)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(createProfile, "", "    ")
	if err != nil {
		log.Printf("CreateProfile - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateProfile - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/profiles/get/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetProfileByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, err := h.service.GetProfileByID(id)
	if err != nil {
		log.Printf("GetProfileByID - h.service.GetProfileByID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Printf("GetProfileByID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var profile model.Profile

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("EditProfile - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditProfile - incoming request: %v", string(data))

	err = json.Unmarshal(data, &profile)
	if err != nil {
		log.Printf("EditProfile - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditProfile - data after unmarshalling: %v", profile)

	updatedProfile, err := h.service.EditProfile(&profile)
	if err != nil {
		log.Printf("EditProfile - h.service.EditProfile error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditProfile - updated profile: %v", updatedProfile)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(updatedProfile, "", "    ")
	if err != nil {
		log.Printf("EditProfile - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditProfile - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/profiles/delete/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteProfile - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.service.DeleteProfile(id)
	if err != nil {
		log.Printf("DeleteProfile - h.service.DeleteProfile error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted profile"))
}
