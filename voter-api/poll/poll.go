package poll

import (
	"encoding/json"
	"errors"
	"fmt"
)

type pollOption struct {
	PollOptionID    uint   `json:"polloptionid"`
	PollOptionValue string `json:"polloptionvalue"`
}

type Poll struct {
	PollID       uint   `json:"pollid"`
	PollTitle    string `json:"polltitle"`
	PollQuestion string `json:"pollquestion"`
	PollOptions  []pollOption
}

type PollMap map[uint]Poll

type PollList struct {
	PollListMap PollMap
	//more things would be included in a real implementation
}

// constructor for VoterList struct

func NewPoll() (*PollList, error) {

	//Now that we know the file exists, at at the minimum we have
	//a valid empty DB, lets create the VoterList struct
	pollList := &PollList{
		PollListMap: make(map[uint]Poll),
	}

	// We should be all set here, the VoterList struct is ready to go
	// so we can support the public database operations
	return pollList, nil
}

func NewPoll1(id uint, title, question string) *Poll {
	return &Poll{
		PollID:       id,
		PollTitle:    title,
		PollQuestion: question,
		PollOptions:  []pollOption{},
	}
}

func NewSamplePoll() *Poll {
	return &Poll{
		PollID:       1,
		PollTitle:    "Favorite Pet",
		PollQuestion: "What type of pet do you like best?",
		PollOptions: []pollOption{
			{PollOptionID: 1, PollOptionValue: "Dog"},
			{PollOptionID: 2, PollOptionValue: "Cat"},
			{PollOptionID: 3, PollOptionValue: "Fish"},
			{PollOptionID: 4, PollOptionValue: "Bird"},
			{PollOptionID: 5, PollOptionValue: "NONE"},
		},
	}
}

func (p *Poll) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (t *PollList) AddPoll(item Poll) error {

	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error
	_, ok := t.PollListMap[uint(item.PollID)]
	if ok {
		return errors.New("item already exists")
	}

	//Now that we know the item doesn't exist, lets add it to our map
	t.PollListMap[uint(item.PollID)] = item
	fmt.Printf("%+v\n", item)
	//If everything is ok, return nil for the error
	return nil
}

func (t *PollList) GetAllPolls() ([]Poll, error) {

	//Now that we have the DB loaded, lets crate a slice
	var pollListList []Poll

	//Now lets iterate over our map and add each item to our slice
	for _, item := range t.PollListMap {
		pollListList = append(pollListList, item)
	}

	//Now that we have all of our items in a slice, return it
	return pollListList, nil
}

func (t *PollList) GetPollItem(id uint) (Poll, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	item, ok := t.PollListMap[id]
	if !ok {
		return Poll{}, errors.New("item does not exist")
	}

	return item, nil
}
