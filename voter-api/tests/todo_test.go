package tests

import (
	"log"
	"os"
	"testing"

	//aliasing package name

	voter "voter-api/voter"

	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// Note the default file path is relative to the test package location.  The
// project has a /tests path where you are at and a /data path where the
// database file sits.  So to get there we need to back up a directory and
// then go into the /data directory.  Thus this is why we are setting the
// default file name to "../data/todo.json"

var (
	BASE_API = "http://localhost:1080"

	cli = resty.New()
)

func newRandToDoItem(id uint) voter.Voter {
	return voter.Voter{
		VoterID:   id,
		FirstName: fake.Person().FirstName,
		LastName:  fake.Person().LastName,
	}
}

func TestVoterList_AddItem(t *testing.T) {

	rsp, err := cli.R().Delete(BASE_API + "/todo")
	if rsp.StatusCode() != 200 {
		log.Printf("error clearing database, %v", err)
		os.Exit(1)
	}

	numLoad := 10
	for i := 1; i <= numLoad; i++ {
		item := newRandToDoItem(uint(i))
		rsp, err := cli.R().
			SetBody(item).
			Post(BASE_API + "/voter")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func TestVoterList_GetAllItems(t *testing.T) {

	var items []voter.Voter

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voter")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, 10, len(items), "Count mateches ")

}

func TestVoterList_GetItem(t *testing.T) {

	var items voter.Voter

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voter/7")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, uint(7), items.VoterID)
	log.Printf("Get voter with ID 7, %v", items)

}
