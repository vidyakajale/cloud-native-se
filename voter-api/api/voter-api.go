package api

import (
	"fmt"

	"log"
	"net/http"
	"os"

	//"strconv"

	"voter-api/poll"
	voter "voter-api/voter"
	election "voter-api/votes"

	"github.com/gofiber/fiber/v2"
)

// The api package creates and maintains a reference to the data handler
// this is a good design practice
type VoterListAPI struct {
	voter *voter.VoterList
}

func NewVoter() (*VoterListAPI, error) {
	dbHandler, err := voter.NewVoter()
	if err != nil {
		return nil, err
	}

	return &VoterListAPI{voter: dbHandler}, nil
}

type PollListAPI struct {
	poll *poll.PollList
}

func NewPoll() (*PollListAPI, error) {
	dbHandler, err := poll.NewPoll()
	if err != nil {
		return nil, err
	}

	return &PollListAPI{poll: dbHandler}, nil
}

type VoteListAPI struct {
	vote *election.VoteList
}

func NewVote() (*VoteListAPI, error) {
	dbHandler, err := election.NewVote()
	if err != nil {
		return nil, err
	}

	return &VoteListAPI{vote: dbHandler}, nil
}

/*func NewPoll() (*VoterListAPI, error) {
	dbHandler, err := poll.NewPoll()
	if err != nil {
		return nil, err
	}

	return &VoterListAPI{poll: dbHandler}, nil
}*/

//Below we implement the API functions.  Some of the framework
//things you will see include:
//   1) How to extract a parameter from the URL, for example
//	  the id parameter in /todo/:id
//   2) How to extract the body of a POST request
//   3) How to return JSON and a correctly formed HTTP status code
//	  for example, 200 for OK, 404 for not found, etc.  This is done
//	  using the c.JSON() function
//   4) How to return an error code and abort the request.  This is
//	  done using the c.AbortWithStatus() function

// implementation for POST /AddVoter
// adds a new Voter
func (td *VoterListAPI) AddVoter(c *fiber.Ctx) error {
	var Voter voter.Voter

	//With HTTP based APIs, a POST request will usually
	//have a body that contains the data to be added
	//to the database.  The body is usually JSON, so
	//we need to bind the JSON to a struct that we
	//can use in our code.
	//This framework exposes the raw body via c.Request.Body
	//but it also provides a helper function BodyParser
	//that will extract the body, convert it to JSON and
	//bind it to a struct for us.  It will also report an error
	//if the body is not JSON or if the JSON does not match
	//the struct we are binding to.

	if err := c.BodyParser(&Voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", Voter)
	if err := td.voter.AddItem(Voter); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", Voter)

	return c.JSON(Voter)
}

// adds a new VoterPoll
func (td *VoterListAPI) AddVoterPoll(c *fiber.Ctx) error {
	var Voter voter.Voter
	var VoterPoll voter.VoterPoll
	id, err := c.ParamsInt("id")
	idUint := uint(id)
	fmt.Printf("idUint %+v\n", idUint)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	pollId, err := c.ParamsInt("pollid")
	pollIdUint := uint(pollId)
	fmt.Printf("pollIdUint %+v\n", pollIdUint)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := c.BodyParser(&VoterPoll); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}
	VoterForPoll, err := td.voter.GetItem(idUint)
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}
	fmt.Printf("VoterForPoll %+v\n", VoterForPoll)
	if err := td.voter.AddVoterPoll(VoterForPoll, VoterPoll.PollID, VoterPoll.VoteID); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}
	Voter = VoterForPoll
	fmt.Printf("%+v\n", Voter)

	return c.JSON(Voter)
}

// implementation for GET /voter
// returns all voters
func (td *VoterListAPI) ListAllVoters(c *fiber.Ctx) error {

	voterList, err := td.voter.GetAllItems()
	if err != nil {
		log.Println("Error Getting All Items: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Items")
	}
	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if voterList == nil {
		voterList = make([]voter.Voter, 0)
	}

	return c.JSON(voterList)
}

// implementation for GET /voter/:id
// returns a single voter
func (td *VoterListAPI) GetVoter(c *fiber.Ctx) error {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id, err := c.ParamsInt("id")
	idUint := uint(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	Voter, err := td.voter.GetItem(idUint)
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	return c.JSON(Voter)
}

// implementation for GET /voter/:id
// returns a single voter
func (td *VoterListAPI) GetVoterPoll(c *fiber.Ctx) error {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id, err := c.ParamsInt("id")
	idUint := uint(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	pollId, err := c.ParamsInt("pollid")
	pollIdUint := uint(pollId)
	fmt.Printf("pollIdUint %+v\n", pollIdUint)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	Voter, err := td.voter.GetPollItem(idUint, pollIdUint)
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	return c.JSON(Voter)
}

//AddPoll

func (td *PollListAPI) AddPoll(c *fiber.Ctx) error {
	var Poll poll.Poll

	if err := c.BodyParser(&Poll); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", Poll)
	if err := td.poll.AddPoll(Poll); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", Poll)

	return c.JSON(Poll)
}

// returns all polls
func (td *PollListAPI) ListAllPolls(c *fiber.Ctx) error {

	pollList, err := td.poll.GetAllPolls()
	if err != nil {
		log.Println("Error Getting All Items: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Items")
	}
	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if pollList == nil {
		pollList = make([]poll.Poll, 0)
	}

	return c.JSON(pollList)
}

// returns a single pole
func (td *PollListAPI) GetPoll(c *fiber.Ctx) error {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id, err := c.ParamsInt("id")
	idUint := uint(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	Poll, err := td.poll.GetPollItem(idUint)
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	return c.JSON(Poll)
}

//All Votes

func (td *VoteListAPI) AddVote(c *fiber.Ctx) error {
	var Vote election.Vote

	if err := c.BodyParser(&Vote); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", Vote)
	if err := td.vote.AddVote(Vote); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", Vote)

	return c.JSON(Vote)
}

// returns all votes
func (td *VoteListAPI) ListAllVotes(c *fiber.Ctx) error {

	voteList, err := td.vote.GetAllVotes()
	if err != nil {
		log.Println("Error Getting All Items: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Items")
	}
	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if voteList == nil {
		voteList = make([]election.Vote, 0)
	}

	return c.JSON(voteList)
}

// returns a single vote
func (td *VoteListAPI) GetVote(c *fiber.Ctx) error {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id, err := c.ParamsInt("id")
	idUint := uint(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	Vote, err := td.vote.GetVoteItem(idUint)
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	return c.JSON(Vote)
}

// implementation for GET /v2/todo
// returns todos that are either done or not done
// depending on the value of the done query parameter
// for example, /v2/todo?done=true will return all
// todos that are done.  Note you can have multiple
// query parameters, for example /v2/todo?done=true&foo=bar
/*
func (td *VoterListAPI) ListSelectTodos(c *fiber.Ctx) error {
	//lets first load the data
	todoList, err := td.voter.GetAllItems()
	if err != nil {
		log.Println("Error Getting Database Items: ", err)
		return fiber.NewError(http.StatusNotFound)
	}
	//If the database is empty, make an empty slice so that the
	//JSON marshalling works correctly
	if todoList == nil {
		todoList = make([]voter.Voter, 0)
	}

	//Note that the query parameter is a string, so we
	//need to convert it to a bool
	doneS := c.Query("done")

	//if the doneS is empty, then we will return all items
	if doneS == "" {
		return c.JSON(todoList)
	}

	//Now we can handle the case where doneS is not empty
	//and we need to filter the list based on the doneS value

	//Now we need to filter the list based on the done value
	//that was passed in.  We will create a new slice and
	//only add items that match the done value
	var filteredList []voter.Voter
	for _, item := range todoList {

		filteredList = append(filteredList, item)

	}

	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if filteredList == nil {
		filteredList = make([]voter.Voter, 0)
	}

	return c.JSON(filteredList)
}
*/

// implementation for PUT /todo
// Web api standards use PUT for Updates
func (td *VoterListAPI) UpdateToDo(c *fiber.Ctx) error {
	var Voter voter.Voter
	if err := c.BodyParser(&Voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := td.voter.UpdateItem(Voter); err != nil {
		log.Println("Error updating item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(Voter)
}

// implementation for DELETE /todo/:id
// deletes a todo
func (td *VoterListAPI) DeleteToDo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	idUint := uint(id)
	if err := td.voter.DeleteItem(idUint); err != nil {
		log.Println("Error deleting item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString("Delete OK")
}

// implementation for DELETE /todo
// deletes all todos
func (td *VoterListAPI) DeleteAllToDo(c *fiber.Ctx) error {

	if err := td.voter.DeleteAll(); err != nil {
		log.Println("Error deleting all items: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString("Delete All OK")
}

/*   SPECIAL HANDLERS FOR DEMONSTRATION - CRASH SIMULATION AND HEALTH CHECK */

// implementation for GET /crash
// This simulates a crash to show some of the benefits of the
// gin framework
func (td *VoterListAPI) CrashSim(c *fiber.Ctx) error {
	//panic() is go's version of throwing an exception
	//note with recover middleware this will not end program
	panic("Simulating an unexpected crash")
}

func (td *VoterListAPI) CrashSim2(c *fiber.Ctx) error {
	//A stupid crash simulation example
	i := 0
	j := 1 / i
	jStr := fmt.Sprintf("%d", j)
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"val_j": jStr,
		})
}

func (td *VoterListAPI) CrashSim3(c *fiber.Ctx) error {
	//A stupid crash simulation example
	os.Exit(10)
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"error": "will never get here, nothing you can do about this",
		})
}

// implementation of GET /health. It is a good practice to build in a
// health check for your API.  Below the results are just hard coded
// but in a real API you can provide detailed information about the
// health of your API with a Health Check
func (td *VoterListAPI) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":             "ok",
			"version":            "1.0.0",
			"uptime":             100,
			"users_processed":    1000,
			"errors_encountered": 10,
		})
}
