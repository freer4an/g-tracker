package models

import (
	"testing"

	"github.com/freer4an/groupie-tracker/internal/helpers"
)

const (
	url = "https://groupietrackers.herokuapp.com/api/artists"
)

func TestArtist_DatesLocations(t *testing.T) {
	artists := getArtists(t)

	tests := []struct {
		name    string
		fields  Artist
		want    int // len(rel.DatesLocations)
		wantErr bool
	}{
		{
			name:   "Queen",
			fields: artists[0],
			want:   8,
		},
		{
			name:   "SOJA",
			fields: artists[1],
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Artist{
				ID:           tt.fields.ID,
				Image:        tt.fields.Image,
				Name:         tt.fields.Name,
				Members:      tt.fields.Members,
				CreationDate: tt.fields.CreationDate,
				FirstAlbum:   tt.fields.FirstAlbum,
				Locations:    tt.fields.Locations,
				Relations:    tt.fields.Relations,
			}
			if a.Name != tt.name {
				t.Errorf("Artist.Name = %v, want %v", a.Name, tt.name)
			}
			got, _ := a.DatesLocations()
			if len(got.DatesLocations) != tt.want {
				t.Errorf("Case %s:\nExpected: %v\nActual: %v", tt.name, tt.want, len(got.DatesLocations))
			}
		})
	}
}

func getArtists(t *testing.T) []Artist {
	var artists []Artist
	if err := helpers.ParseAPI(url, &artists); err != nil {
		t.Errorf("Parse error: %v", err)
		t.FailNow()
	}
	return artists
}
