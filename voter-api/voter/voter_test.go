package db

import (
	"reflect"
	"testing"
)

func TestVoterList_GetAllItems(t *testing.T) {
	type fields struct {
		VoterListMap VoterMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &VoterList{
				VoterListMap: tt.fields.VoterListMap,
			}
			got, err := tr.GetAllItems()
			if (err != nil) != tt.wantErr {
				t.Errorf("VoterList.GetAllItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoterList.GetAllItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
