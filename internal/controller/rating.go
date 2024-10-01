package controller

import (
	"context"
	"errors"
	"metadata.com/internal/reporitory"
	"metadata.com/model"
)

var ErrNotFound = errors.New("rating not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordType model.RecordType, recordID model.RecordID) ([]model.Rating, error)
	Put(ctx context.Context, recordType model.RecordType, recordID model.RecordID, rating *model.Rating) error
}

type Controller struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (r *Controller) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := r.repo.Get(ctx, recordType, recordID)
	if err != nil && errors.Is(err, reporitory.ErrNotFound) {
		return 0, ErrNotFound
	}

	sum := float64(0)
	for _, val := range ratings {
		sum += float64(val.Value)
	}

	avg := sum / float64(len(ratings))
	return avg, nil
}

func (r *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return r.repo.Put(ctx, recordType, recordID, rating)
}
