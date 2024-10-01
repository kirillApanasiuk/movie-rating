package model

type RecordID string

// Specify specific record type
type RecordType string

const (
	RecordTypeMovie RecordType = "microservices_in_go"
)

type UserId string
type RatingValue int

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserId      `json:"userId"`
	Value      RatingValue `json:"value"`
}
