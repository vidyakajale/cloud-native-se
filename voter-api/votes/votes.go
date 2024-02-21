package election

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Vote struct {
	VoteID    uint `json:"voteid"`
	VoterID   uint `json:"voterid"`
	PollID    uint `json:"pollid"`
	VoteValue uint `json:"votevalue"`
}

type VoteMap map[uint]Vote

type VoteList struct {
	VoteListMap VoteMap
	//more things would be included in a real implementation
}

// constructor for VoterList struct

func NewVote() (*VoteList, error) {

	//Now that we know the file exists, at at the minimum we have
	//a valid empty DB, lets create the VoterList struct
	voteList := &VoteList{
		VoteListMap: make(map[uint]Vote),
	}

	// We should be all set here, the VoterList struct is ready to go
	// so we can support the public database operations
	return voteList, nil
}
func (t *VoteList) AddVote(item Vote) error {

	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error
	_, ok := t.VoteListMap[uint(item.VoteID)]
	if ok {
		return errors.New("item already exists")
	}

	//Now that we know the item doesn't exist, lets add it to our map
	t.VoteListMap[uint(item.VoteID)] = item
	VoterId := item.VoterID
	fmt.Printf("VoterId = %+v\n", VoterId)

	fmt.Printf("%+v\n", item)
	//If everything is ok, return nil for the error
	return nil
}

func (t *VoteList) GetAllVotes() ([]Vote, error) {

	//Now that we have the DB loaded, lets crate a slice
	var voteListList []Vote

	//Now lets iterate over our map and add each item to our slice
	for _, item := range t.VoteListMap {
		voteListList = append(voteListList, item)
	}

	//Now that we have all of our items in a slice, return it
	return voteListList, nil
}

func (t *VoteList) GetVoteItem(id uint) (Vote, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	item, ok := t.VoteListMap[id]
	if !ok {
		return Vote{}, errors.New("item does not exist")
	}

	return item, nil
}

// constructor for VoterList struct
func NewVote1(pid, vid, vtrid, vval uint) *Vote {
	return &Vote{
		VoteID:    vid,
		VoterID:   vtrid,
		PollID:    pid,
		VoteValue: vval,
	}
}

func NewSampleVote() *Vote {
	return &Vote{
		VoteID:    1,
		PollID:    1,
		VoterID:   1,
		VoteValue: 1,
	}
}

func (p *Vote) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
