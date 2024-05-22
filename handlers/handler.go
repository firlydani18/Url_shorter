package handlers

import (
	"backend/models"
	"backend/repositories"
	"backend/utils/respons"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type handler struct {
	LinkRepository repositories.LinkRepository
}

func HandlerLink(LinkRepository repositories.LinkRepository) *handler {
	return &handler{LinkRepository}
}

func (h *handler) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["unique_id"]
	user, err := h.LinkRepository.GetLink(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := respons.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get(user.LongURL)
	if err != nil {
		http.Redirect(w, r, "http://"+user.LongURL, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, user.LongURL, http.StatusSeeOther)
	}

	fmt.Println(resp)

	w.WriteHeader(http.StatusOK)
	response := respons.SuccessResult{StatusCode: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	shortURL := models.Link{
		ShortURL: r.FormValue("short_url"),
	}

	if err := json.NewDecoder(r.Body).Decode(&shortURL); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := respons.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ := h.LinkRepository.GetLink(shortURL.ShortURL[22:])

	linkData := models.Link{
		ID:       data.ID,
		ShortURL: data.ShortURL,
		LongURL:  data.LongURL,
	}

	w.WriteHeader(http.StatusOK)
	response := respons.SuccessResult{StatusCode: http.StatusOK, Data: responseLink(linkData)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	longURL := models.Link{
		LongURL: r.FormValue("long_url"),
	}

	if err := json.NewDecoder(r.Body).Decode(&longURL); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := respons.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err := h.LinkRepository.GetLongURLCheck(longURL.LongURL)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := respons.ErrorResult{StatusCode: http.StatusBadRequest, Message: "URL already shorten! If you don't remember it, please use get original URL service!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	rand.Seed(time.Now().UnixNano())
	shortURL := randSeq(7)

	linkData := models.Link{
		ShortURL: shortURL,
		LongURL:  longURL.LongURL,
	}

	linkData, _ = h.LinkRepository.CreateLink(linkData)

	w.WriteHeader(http.StatusOK)
	response := respons.SuccessResult{StatusCode: http.StatusOK, Data: responseLink(linkData)}
	json.NewEncoder(w).Encode(response)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func responseLink(p models.Link) models.Link {
	return models.Link{
		ID:       p.ID,
		ShortURL: "http://localhost:5000/" + p.ShortURL,
		LongURL:  p.LongURL,
	}
}
