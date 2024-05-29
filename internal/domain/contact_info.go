package domain

type ContactInfo struct {
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
}
