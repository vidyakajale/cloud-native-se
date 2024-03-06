package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
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

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:55013"
	RedisKeyPrefix       = "todo:"
)

type cache struct {
	client  *redis.Client
	context context.Context
}

// VoterMap is a type alias for a map of Voters.  The key
// will be the Voter.Id and the value will be the Voter
//type VoterMap map[uint]Voter

// VoterList is the struct that represents the main object of our
// VoterList app.  It contains a map of Voters and the name of
// the file that is used to store the items.
//
// This is just a mock, so we will only be managing an in memory
// map
type VoterList struct {
	//VoterListMap VoterMap
	cache
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
	//voterList := &VoterList{
	//		VoterListMap: make(map[uint]Voter),
	//	}

	// We should be all set here, the VoterList struct is ready to go
	// so we can support the public database operations
	//	return voterList, nil

	//We will use an override if the REDIS_URL is provided as an environment
	//variable, which is the preferred way to wire up a docker container
	redisUrl := os.Getenv("REDIS_URL")
	//This handles the default condition
	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}
	return NewWithCacheInstance(redisUrl)
}

func NewWithCacheInstance(location string) (*VoterList, error) {

	//Connect to redis.  Other options can be provided, but the
	//defaults are OK
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	//We use this context to coordinate betwen our go code and
	//the redis operaitons
	ctx := context.TODO()

	//This is the reccomended way to ensure that our redis connection
	//is working
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error())
		return nil, err
	}

	//Return a pointer to a new ToDo struct
	return &VoterList{
		cache: cache{
			client:  client,
			context: ctx,
		},
	}, nil
}

//------------------------------------------------------------
// REDIS HELPERS
//------------------------------------------------------------

// We will use this later, you can ignore for now
/*
func isRedisNilError(err error) bool {
	return errors.Is(err, redis.Nil) || err.Error() == RedisNilError
}*/

// In redis, our keys will be strings, they will look like
// todo:<number>.  This function will take an integer and
// return a string that can be used as a key in redis
func redisKeyFromId(id int) string {
	return fmt.Sprintf("%s%d", RedisKeyPrefix, id)
}

// getAllKeys will return all keys in the database that match the prefix
// used in this application - RedisKeyPrefix.  It will return a string slice
// of all keys.  Used by GetAll and DeleteAll
func (t *VoterList) getAllKeys() ([]string, error) {
	key := fmt.Sprintf("%s*", RedisKeyPrefix)
	return t.client.Keys(t.context, key).Result()
}

func fromJsonString(s string, item *Voter) error {
	err := json.Unmarshal([]byte(s), &item)
	if err != nil {
		return err
	}
	return nil
}

// upsertToDo will be used by insert and update, Redis only supports upserts
// so we will check if an item exists before update, and if it does not exist
// before insert
func (t *VoterList) upsertToDo(item *Voter) error {
	log.Println("Adding new Id:", redisKeyFromId(int(item.VoterID)))
	return t.client.JSONSet(t.context, redisKeyFromId(int(item.VoterID)), ".", item).Err()
}

// Helper to return a ToDoItem from redis provided a key
func (t *VoterList) getItemFromRedis(key string, item *Voter) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemJson, err := t.client.JSONGet(t.context, key, ".").Result()
	if err != nil {
		return err
	}

	return fromJsonString(itemJson, item)
}

func (t *VoterList) doesKeyExist(id int) bool {
	kc, _ := t.client.Exists(t.context, redisKeyFromId(id)).Result()
	return kc > 0
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
func (t *VoterList) AddItem(item *Voter) error {

	if t.doesKeyExist(int(item.VoterID)) {
		return fmt.Errorf("ToDo item with id %d already exists", item.VoterID)
	}
	return t.upsertToDo(item)
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
func (t *VoterList) DeleteItem(id int) error {

	if !t.doesKeyExist(id) {
		return fmt.Errorf("ToDo item with id %d does not exist", id)
	}
	return t.client.Del(t.context, redisKeyFromId(id)).Err()
}

// DeleteAll removes all items from the DB.
// It will be exposed via a DELETE /todo endpoint
func (t *VoterList) DeleteAll() (int, error) {
	keyList, err := t.getAllKeys()
	if err != nil {
		return 0, err
	}

	//Notice how we can deconstruct the slice into a variadic argument
	//for the Del function by using the ... operator
	numDeleted, err := t.client.Del(t.context, keyList...).Result()
	return int(numDeleted), err

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
func (t *VoterList) UpdateItem(item *Voter) error {

	if !t.doesKeyExist(int(item.VoterID)) {
		return fmt.Errorf("Voter with id %d does not exist", item.VoterID)
	}
	return t.upsertToDo(item)
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
func (t *VoterList) GetItem(id int) (*Voter, error) {

	newToDo := &Voter{}
	err := t.getItemFromRedis(redisKeyFromId(id), newToDo)
	if err != nil {
		return nil, err
	}
	return newToDo, nil
}

func (t *VoterList) GetPollItem(id uint, pollId uint) (Voter, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	var itemForPoll Voter
	/*item, ok := t.VoterListMap[id]
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
	}*/
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

	keyList, err := t.getAllKeys()
	if err != nil {
		return nil, err
	}

	//preallocate the slice, will make things faster
	resList := make([]Voter, len(keyList))

	for idx, k := range keyList {
		err := t.getItemFromRedis(k, &resList[idx])
		if err != nil {
			return nil, err
		}
	}

	return resList, nil

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

func (t *VoterList) AddVoterPoll(item *Voter, pollID uint, voteID uint) error {

	v := item
	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error

	v.AddPollToVoter(pollID, voteID)

	//Now that we know the item doesn't exist, lets add it to our map

	//t.VoterListMap[uint(item.VoterID)] = v
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
