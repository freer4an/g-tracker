package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

const (
	apiURL      = "https://groupietrackers.herokuapp.com/api"
	artistURL   = "/artists"
	relationURL = "/relation"
)

var (
	apiData *api
	ticker  = time.Tick(10 * time.Second)
)

type api struct {
	Bands     []Band
	Relations []Relation
	len       int
	mu        sync.Mutex
}

func (d *api) CheckID(index int) bool {
	return index >= 0 && index < d.len
}

func (data *api) fill() error {
	data.mu.Lock()
	defer data.mu.Unlock()
	var rel Relations
	if err := helpers.ParseAPI(apiURL+artistURL, &data.Bands); err != nil {
		return fmt.Errorf("Bands parse error: %w", err)
	}
	if err := helpers.ParseAPI(apiURL+relationURL, &rel); err != nil {
		return fmt.Errorf("Relations parse error: %w", err)
	}
	data.Relations = rel.Index
	data.len = len(data.Bands)
	return nil
}

func updateChecker() {
	if apiData == nil {
		return
	}
	data := new(api)
	if err := helpers.ParseAPI(apiURL+artistURL, &data.Bands); err != nil {
		return
	}
	if data.len == apiData.len {
		return
	}
	apiData.fill()
}
