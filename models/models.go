package models

import (
	"time"

	"github.com/google/uuid"
)

type Coffee struct {
	Name     string
	Notes    string
	Duration int
	In       int
	Out      int
	Date     time.Time
	Rating   int
	Grind    int
	GUID     string
}

func NewCoffee(name string, notes string, duration int, in int, out int, rating int, grind int) *Coffee {
	return &Coffee{
		Name:     name,
		Notes:    notes,
		Duration: duration,
		In:       in,
		Out:      out,
		Date:     time.Now(),
		Rating:   rating,
		Grind:    grind,
		GUID:     generateGUID(),
	}
}

func generateGUID() string {
	id := uuid.New()
	return id.String()
}
