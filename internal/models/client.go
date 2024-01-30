package models

import (
	"errors"
	"fmt"
)

// errors
var (
	ErrID = errors.New("Invalid id")
)

type clientData struct {
	API      *api
	Band     Band
	Relation Relation
}

func GetHomeData() (*clientData, error) {
	if apiData == nil {
		apiData = new(api)
		if err := apiData.fill(); err != nil {
			return nil, err
		}
	}

	data := &clientData{API: apiData}

	select {
	case <-ticker:
		go updateChecker()
	default:
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
		return nil, fmt.Errorf("%w - %d", ErrID, id)
	}
	return &clientData{
		Band:     apiData.Bands[id],
		Relation: apiData.Relations[id],
	}, nil
}
