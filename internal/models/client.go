package models

import (
	"fmt"
	"time"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

// var ticker = make(chan time.Time)
var ticker = time.Tick(10 * time.Second)

type clientData struct {
	API       *api
	ArtistURL string
	Band      Band
	Relation  Relation
}

func GetHomeData() (*clientData, error) {
	defer func() {
		go updateChecker()
	}()
	if apiData == nil {
		apiData = new(api)
		if err := apiData.fill(); err != nil {
			return nil, err
		}
	}

	data := &clientData{
		API:       apiData,
		ArtistURL: artistURL,
	}

	return data, nil
}

func GetBandData(id int) (*clientData, error) {
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
	return &clientData{
		Band:     apiData.Bands[id],
		Relation: apiData.Relations[id],
	}, nil
}

func updateChecker() {
	select {
	case <-ticker:
		if apiData == nil {
			return
		}
		data := new(api)
		if err := helpers.ParseAPI(apiURL+artistURL, &data.Bands); err != nil {
			return
		}
		if data.len != apiData.len {
			apiData.mu.Lock()
			defer apiData.mu.Unlock()
			if err := apiData.fill(); err != nil {
				return
			}
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
