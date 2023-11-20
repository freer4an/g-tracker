package models

import "sort"

type Relation struct {
	// DatesLocations map[string][]string `json:"datesLocations"`
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
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
