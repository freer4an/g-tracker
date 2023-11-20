package models

import (
	"reflect"
	"testing"
)

func TestRelation_Locations(t *testing.T) {
	artists := getArtists(t)
	queen, err := artists[0].DatesLocations()
	if err != nil {
		t.Errorf("Relation parse error: %v", err)
		t.FailNow()
	}

	tests := []struct {
		name   string
		fields Relation
		want   []string
	}{
		{
			name:   "Queen",
			fields: *queen,
			want: []string{"dunedin-new_zealand", "georgia-usa",
				"los_angeles-usa", "nagoya-japan", "north_carolina-usa",
				"osaka-japan", "penrose-new_zealand", "saitama-japan",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Relation{
				ID:             tt.fields.ID,
				DatesLocations: tt.fields.DatesLocations,
			}
			if got := r.Locations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Relation.Locations() = %v, want %v", got, tt.want)
			}
		})
	}
}
