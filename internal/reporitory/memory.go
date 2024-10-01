package reporitory

import (
	"context"
	"errors"
	"metadata.com/model"
)

var ErrNotFound = errors.New("reporitory not found")

type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

func NewRepository() *Repository {
	return &Repository{
		map[model.RecordType]map[model.RecordID][]model.Rating{},
	}
}

func (r *Repository) Get(_ context.Context, recordType model.RecordType, id model.RecordID) ([]model.Rating, error) {
	recordsForType, ok := r.data[recordType]
	if !ok {
		return nil, ErrNotFound
	}

	ratings, ok := recordsForType[id]
	if !ok || len(ratings) == 0 {
		return nil, ErrNotFound
	}

	return ratings, nil
}

func (r *Repository) Put(_ context.Context, recordType model.RecordType, id model.RecordID, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}

	r.data[recordType][id] = append(r.data[recordType][id], *rating)

	return nil
}
