package main

import (
	"github.com/kirillApanasiuk/movie-rating/internal/controller"
	ratingHttp "github.com/kirillApanasiuk/movie-rating/internal/handler/http"
	"github.com/kirillApanasiuk/movie-rating/internal/reporitory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := reporitory.NewRepository()
	ctrl := controller.New(repo)
	h := ratingHttp.New(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
