package models

import (
	"fmt"
	"time"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

// var ticker = make(chan time.Time)
var ticker = time.Tick(10 * time.Second)

type payloadData struct {
	*api
	ArtistURL      string
	Artist         *Artist
	DatesLocations *Relation
}

func GetHomeData() (*payloadData, error) {
	defer func() {
		go updateChecker()
	}()
	if apiData == nil {
		apiData = new(api)
		if err := apiData.fill(); err != nil {
			return nil, err
		}
	}

	data := &payloadData{
		api:       apiData,
		ArtistURL: artistURL,
	}

	return data, nil
}

func GetArtistData(id int) (*payloadData, error) {
	if apiData == nil {
		apiData = new(api)
		if err := apiData.fill(); err != nil {
			return nil, err
		}
	}
	id = id - 1
	if !apiData.CheckID(id) {
		return nil, fmt.Errorf("Invalid id: %d", id)
	}
	return &payloadData{
		Artist:         &apiData.Artists[id],
		DatesLocations: &apiData.Relation[id],
	}, nil
}

func updateChecker() {
	select {
	case <-ticker:
		if apiData == nil {
			return
		}
		data := new(api)
		if err := helpers.ParseAPI(apiURL+artistURL, &data.Artists); err != nil {
			return
		}
		if data.len != apiData.len {
			apiData.mu.Lock()
			if err := apiData.fill(); err != nil {
				return
			}
			apiData.mu.Unlock()
		}
	default:
		return
	}
}

// func init() {
// 	go func() {
// 		var ticker = time.Tick(10 * time.Second)
// 	}()
// }
