package interfaces

import "time"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    Gender `json:"gender"`
}

type Couple struct {
	Person1 Person `json:"person_1"`
	Person2 Person `json:"person_2"`
}

type Wedding struct {
	ID         string    `json:"id"`
	Couple     Couple    `json:"couple"`
	Date       time.Time `json:"date"`
	Location   string    `json:"location"`
	GuestCount int       `json:"guest_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
