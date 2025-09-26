package main

import (
	"io"
	"net/http"
	"strings"

	spinhttp "github.com/spinframework/spin-go-sdk/v2/http"
	"github.com/spinframework/spin-go-sdk/v2/kv"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()

		router.GET("/:slug", redirect)
		router.POST("/:slug", shorten)

		router.ServeHTTP(w, r)
	})
}

func redirect(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	slug := params.ByName("slug")
	store, err := kv.OpenStore("default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer store.Close()

	url, err := store.Get(slug)
	if err != nil {
		if strings.Contains(err.Error(), "no such key") {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, string(url), http.StatusFound)
}

func shorten(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	slug := params.ByName("slug")
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store, err := kv.OpenStore("default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer store.Close()

	err = store.Set(slug, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
