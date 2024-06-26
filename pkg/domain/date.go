package domain

// Date represents a date with day, month, and year.
type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}

// TimeRange represents a time range with start and end times. Useful
// for specifying the time range of games.
type TimeRange struct {
	Start TimeOfDay `json:"start" bson:"start"`
	End   TimeOfDay `json:"end" bson:"end"`
}

// TimeOfDay represents a time of day with hour and minute.
type TimeOfDay struct {
	Hour   int `json:"hour" bson:"hour"`
	Minute int `json:"minute" bson:"minute"`
}
