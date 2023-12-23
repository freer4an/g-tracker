package models

import "sort"

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation
}

func (r *Relation) Locations() []string {
	var locations []string
	for k := range r.DatesLocations {
		locations = append(locations, k)
	}
	sort.Slice(locations, func(i, j int) bool {
		return locations[i] < locations[j]
	})
	return locations
}
