package domain

// ContactInfo represents the contact information for a player.
type ContactInfo struct {
	Email string `json:"email" bson:"email" binding:"email"`
	Phone string `json:"phone" bson:"phone"`
}
