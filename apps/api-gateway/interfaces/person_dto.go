package interfaces

type GenderDTO string

const (
	Male   GenderDTO = "male"
	Female GenderDTO = "female"
	Other  GenderDTO = "other"
)

type PersonDTO struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    GenderDTO `json:"gender"`
}
