package models

import (
	"fmt"
	"sync"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

const (
	apiURL      = "https://groupietrackers.herokuapp.com/api"
	artistURL   = "/artists"
	relationURL = "/relation"
)

var (
	apiData *api
)

type api struct {
	Artists  []Artist
	Relation []Relation
	len      int
	mu       sync.Mutex
}

func (d *api) CheckID(index int) bool {
	return index >= 0 && index < d.len
}

func (data *api) fill() error {
	data.mu.Lock()
	defer data.mu.Unlock()
	if data.Artists != nil && data.Relation != nil {
		return nil
	}

	var rel Relations
	if err := helpers.ParseAPI(apiURL+artistURL, &data.Artists); err != nil {
		return fmt.Errorf("Artists parse error: %w", err)
	}
	if err := helpers.ParseAPI(apiURL+relationURL, &rel); err != nil {
		return fmt.Errorf("Relation parse error: %w", err)
	}
	data.Relation = rel.Index
	data.len = len(data.Artists)
	return nil
}
