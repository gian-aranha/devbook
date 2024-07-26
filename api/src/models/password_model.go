package models

// Password represents the update password request format
type Passord struct {
	New string `json:"new"`
	Current string `json:"current"`
}