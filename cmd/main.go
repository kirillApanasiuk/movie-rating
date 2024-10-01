package main

import (
	"log"
	"metadata.com/internal/controller"
	ratingHttp "metadata.com/internal/handler/http"
	"metadata.com/internal/reporitory"
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
