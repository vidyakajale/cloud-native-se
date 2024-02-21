package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// voterPoll is the struct that represent a single poll
type VoterPoll struct {
	PollID   uint      `json:"pollid"`
	VoteID   uint      `json:"voteid"`
	VoteDate time.Time `json:"votedate"`
}

// Voter is the struct that represents a single voter item
type Voter struct {
	VoterID     uint   `json:"voterid"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	VoteHistory []VoterPoll
	//IsDone      bool `json:"done"`
}

// VoterMap is a type alias for a map of Voters.  The key
// will be the Voter.Id and the value will be the Voter
type VoterMap map[uint]Voter

// VoterList is the struct that represents the main object of our
// VoterList app.  It contains a map of Voters and the name of
// the file that is used to store the items.
//
// This is just a mock, so we will only be managing an in memory
// map
type VoterList struct {
	VoterListMap VoterMap
	//more things would be included in a real implementation
}

// New is a constructor function that returns a pointer to a new
// VoterList struct.  It takes a single string argument that is the
// name of the file that will be used to store the VoterList items.
// If the file doesn't exist, it will be created.  If the file
// does exist, it will be loaded into the VoterList struct.
func NewVoter() (*VoterList, error) {

	//Now that we know the file exists, at at the minimum we have
	//a valid empty DB, lets create the VoterList struct
	voterList := &VoterList{
		VoterListMap: make(map[uint]Voter),
	}

	// We should be all set here, the VoterList struct is ready to go
	// so we can support the public database operations
	return voterList, nil
}

//------------------------------------------------------------
// THESE ARE THE PUBLIC FUNCTIONS THAT SUPPORT OUR TODO APP
//------------------------------------------------------------

// AddItem accepts a Voter and adds it to the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must not already exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if so, return an error
//
// Postconditions:
//
//	 (1) The item will be added to the DB
//		(2) The DB file will be saved with the item added
//		(3) If there is an error, it will be returned
func (t *VoterList) AddItem(item Voter) error {

	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error
	_, ok := t.VoterListMap[uint(item.VoterID)]
	if ok {
		return errors.New("item already exists")
	}

	//Now that we know the item doesn't exist, lets add it to our map
	t.VoterListMap[uint(item.VoterID)] = item
	fmt.Printf("%+v\n", item)
	//If everything is ok, return nil for the error
	return nil
}

// DeleteItem accepts an item id and removes it from the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be removed from the DB
//		(2) The DB file will be saved with the item removed
//		(3) If there is an error, it will be returned
func (t *VoterList) DeleteItem(id uint) error {

	// we should if item exists before trying to delete it
	// this is a good practice, return an error if the
	// item does not exist

	//Now lets use the built-in go delete() function to remove
	//the item from our map
	delete(t.VoterListMap, id)

	return nil
}

// DeleteAll removes all items from the DB.
// It will be exposed via a DELETE /todo endpoint
func (t *VoterList) DeleteAll() error {
	//To delete everything, we can just create a new map
	//and assign it to our existing map.  The garbage collector
	//will clean up the old map for us
	t.VoterListMap = make(map[uint]Voter)

	return nil
}

// UpdateItem accepts a Voter and updates it in the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be updated in the DB
//		(2) The DB file will be saved with the item updated
//		(3) If there is an error, it will be returned
func (t *VoterList) UpdateItem(item Voter) error {

	// Check if item exists before trying to update it
	// this is a good practice, return an error if the
	// item does not exist
	_, ok := t.VoterListMap[uint(item.VoterID)]
	if !ok {
		return errors.New("item does not exist")
	}

	//Now that we know the item exists, lets update it
	t.VoterListMap[uint(item.VoterID)] = item

	return nil
}

// GetItem accepts an item id and returns the item from the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be returned, if it exists
//		(2) If there is an error, it will be returned
//			along with an empty Voter
//		(3) The database file will not be modified
func (t *VoterList) GetItem(id uint) (Voter, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	item, ok := t.VoterListMap[id]
	if !ok {
		return Voter{}, errors.New("item does not exist")
	}

	return item, nil
}

func (t *VoterList) GetPollItem(id uint, pollId uint) (Voter, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	var itemForPoll Voter
	item, ok := t.VoterListMap[id]
	if !ok {
		return Voter{}, errors.New("item does not exist")
	}
	itemForPoll.FirstName = item.FirstName
	itemForPoll.VoterID = item.VoterID
	itemForPoll.LastName = item.LastName
	for _, val := range item.VoteHistory {
		fmt.Printf("Voter value%+v\n", val)
		if val.PollID == pollId {
			itemForPoll.VoteHistory = append(itemForPoll.VoteHistory, val)
		}
	}
	return itemForPoll, nil
}

// ChangeItemDoneStatus accepts an item id and a boolean status.
// It returns an error if the status could not be updated for any
// reason.  For example, the item itself does not exist, or an
// IO error trying to save the updated status.

// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The items status in the database will be updated
//		(2) If there is an error, it will be returned.
//		(3) This function MUST use existing functionality for most of its
//			work.  For example, it should call GetItem() to get the item
//			from the DB, then it should call UpdateItem() to update the
//			item in the DB (after the status is changed).
func (t *VoterList) ChangeItemDoneStatus(id int, value bool) error {

	//update was successful
	return errors.New("not implemented")
}

// GetAllItems returns all items from the DB.  If successful it
// returns a slice of all of the items to the caller
// Preconditions:   (1) The database file must exist and be a valid
//
// Postconditions:
//
//	 (1) All items will be returned, if any exist
//		(2) If there is an error, it will be returned
//			along with an empty slice
//		(3) The database file will not be modified
func (t *VoterList) GetAllItems() ([]Voter, error) {

	//Now that we have the DB loaded, lets crate a slice
	var voterListList []Voter

	//Now lets iterate over our map and add each item to our slice
	for _, item := range t.VoterListMap {
		voterListList = append(voterListList, item)
	}

	//Now that we have all of our items in a slice, return it
	return voterListList, nil
}

// PrintItem accepts a Voter and prints it to the console
// in a JSON pretty format. As some help, look at the
// json.MarshalIndent() function from our in class go tutorial.
func (t *VoterList) PrintItem(item Voter) {
	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
	fmt.Println(string(jsonBytes))
}

// PrintAllItems accepts a slice of Voters and prints them to the console
// in a JSON pretty format.  It should call PrintItem() to print each item
// versus repeating the code.
func (t *VoterList) PrintAllItems(itemList []Voter) {
	for _, item := range itemList {
		t.PrintItem(item)
	}
}

func (t *VoterList) AddVoterPoll(item Voter, pollID uint, voteID uint) error {

	v := item
	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error

	v.AddPollToVoter(pollID, voteID)
	//Voter.VoteHistory=append([]VoteHistory , VoterPoll{PollID: pollID, VoteID: voteID, VoteDate: time.Now()})

	//Now that we know the item doesn't exist, lets add it to our map
	//t.VoterListMap[uint(item.VoterID)] = item
	t.VoterListMap[uint(item.VoterID)] = v
	fmt.Printf("%+v\n", v)
	//If everything is ok, return nil for the error
	return nil
}

func (v *Voter) AddPollToVoter(pollID uint, voteID uint) {
	v.VoteHistory = append(v.VoteHistory, VoterPoll{PollID: pollID, VoteID: voteID, VoteDate: time.Now()})

}

func (v *Voter) AddVoterPollWithTimeDetails(pollID uint, timeOfPoll time.Time) {
	v.VoteHistory = append(v.VoteHistory, VoterPoll{PollID: pollID, VoteDate: timeOfPoll})

}

// JsonToItem accepts a json string and returns a Voter
// This is helpful because the CLI accepts todo items for insertion
// and updates in JSON format.  We need to convert it to a Voter
// struct to perform any operations on it.
func (t *VoterList) JsonToItem(jsonString string) (Voter, error) {
	var item Voter
	err := json.Unmarshal([]byte(jsonString), &item)
	if err != nil {
		return Voter{}, err
	}
	fmt.Printf("%+v\n", item)
	return item, nil
}
