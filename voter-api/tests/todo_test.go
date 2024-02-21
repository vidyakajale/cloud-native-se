package tests

import (
	"reflect"
	"testing"

	//aliasing package name
	//voter "voter-api/voter"
	voter "voter-api/voter"

	"github.com/stretchr/testify/assert"
	//"github.com/go-resty/resty/v2"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

// Note the default file path is relative to the test package location.  The
// project has a /tests path where you are at and a /data path where the
// database file sits.  So to get there we need to back up a directory and
// then go into the /data directory.  Thus this is why we are setting the
// default file name to "../data/todo.json"
/*
var (
	BASE_API = "http://localhost:1080"

	cli = resty.New()
)*/

func TestVoterList_GetAllItems(t *testing.T) {
	type fields struct {
		VoterListMap voter.VoterMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    []voter.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &voter.VoterList{
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
			assert.NoError(t, err, "Error getting item from Voterlist")
		})
	}
	t.Log("Testing Get All Item: ")

	//Use GetAllItem to get list

	count := 1

	assert.Equal(t, count, 1, "Count mateches ")
}

/*
func TestNewVoter(t *testing.T) {
	tests := []struct {
		name    string
		want    *VoterListAPI
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVoter()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVoter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVoter(t *testing.T) {
	VoterList := [] struct {

	}
	tests := []struct {
		name    string
		want    *VoterList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVoter()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVoter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterListAPI_AddVoter(t *testing.T) {
	type fields struct {
		voter *voter.VoterList
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	} {
	    var items []voter.ToDoItem

		rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voter")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())

		assert.Equal(t, 3, len(items))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &VoterListAPI{
				voter: tt.fields.voter,
			}
			if err := td.AddVoter(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("VoterListAPI.AddVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/
