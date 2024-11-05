package models

type Observer interface {
	Update(status string)
	GetID() int // New method to identify observer by ID
}
