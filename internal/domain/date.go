package domain

// Date represents a date with day, month, and year.
type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}
