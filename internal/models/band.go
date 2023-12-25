package models

import (
	"fmt"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
}

func (a *Band) DatesLocations() (*Relation, error) {
	rel := &Relation{}
	if a.Relations == "" {
		return rel, nil
	}
	if err := helpers.ParseAPI(a.Relations, &rel); err != nil {
		return nil, fmt.Errorf("Parse error: %w", err)
	}
	return rel, nil
}
