package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
)

type FilmController struct {
    store *store.FilmStorer
}

func NewFilmController(store *store.FilmStorer) *FilmController {
    return &FilmController{
        store: store,
    }
}

func (f *FilmController) GetFilms(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodGet {
        id := r.PathValue("id")

        i, err := strconv.Atoi(id)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        resp, err := f.store.GetFilmByUser(i)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        
        body, _ := json.Marshal(resp)
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write(body)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func (f *FilmController) AddFilm(w http.ResponseWriter, r *http.Request) {
    user_id := r.PathValue("id")

    user_id_checked, err := strconv.Atoi(user_id)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    rate := r.FormValue("rate")

    rate_checked, err := strconv.Atoi(rate) 
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    film := models.Film{
        Name: r.FormValue("name"),
        Genre: r.FormValue("genre"),
        Review: r.FormValue("review"),
        Rate: int64(rate_checked),
        UserID: int64(user_id_checked),
    }

    _, err = f.store.AddFilmToUser(film)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    redirectUrl := "/user/%d/film/add"

    fmt.Printf(redirectUrl, user_id_checked)

    http.Redirect(w, r, redirectUrl, http.StatusFound)

}

