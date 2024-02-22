## Voter API

This is a describes application showing API's for voter-api


**Folders:**
    api   :    voter-api-handler
    voters: 
    poll  :
    votes :

**API's:**

   **ListAllVoters:** localhost:1080/voter  (GET)
	   app.Get("/voter", apiHandler.ListAllVoters)  
       Testcases: Passed

   
   
   **Add Voter:**  localhost:1080/voter   (POST)
    	app.Post("/voter", apiHandler.AddVoter) 
        Test cases :Passed
        Payload:
                 {
                    "voterid": 14,
                    "firstname": "William",
                    "lastname": "Torning",
                    "VoteHistory": [{"pollid": 1,"voteid": 1,"VoteDate":"2023-07-25T19:10:34.811997-04:00"},{"pollid": 2,"voteid": 2,"VoteDate":"2023-08-25T19:10:34.811997-04:00"}]
                }

    
    **Get Voter by id:**  localhost:1080/voter/{id}   (GET)
        app.Get("/voter/:id<uint>", apiHandler.GetVoter) 
        Test case: Passed
	
 
 
     **Add Voter by id:**  localhost:1080/voter{id}/poll/{pollid}  (POST)
         app.Post("/voter/:id<uint>/poll/:pollid<uint>", apiHandler.AddVoterPoll)
         Paylpad:
                {
                        "pollid": 2,
                        "voteid": 45
                }

 
	app.Get("/voter/:id<uint>/poll/:pollid<uint>", apiHandler.GetVoterPoll)
	app.Get("/voters/health", apiHandler.HealthCheck)
	app.Post("/poll", apiHandler1.AddPoll)
{
    "pollid": 1,
    "polltitle": "Favorite Pet",
    "pollquestion": "What type of pet do you like best?",
    "PollOptions": [{"polloptionid": 1,"polloptionvalue":"cat"},{"polloptionid": 2,"polloptionvalue":"dog"}]
}

{
    "pollid": 2,
    "polltitle": "Favorite coke",
    "pollquestion": "What type of coke do you like best?",
    "PollOptions": [{"polloptionid": 1,"polloptionvalue":"pepsi"},{"polloptionid": 2,"polloptionvalue":"fanta"}]
}

 
	app.Get("/poll", apiHandler1.ListAllPolls)
	app.Get("/poll/:id<uint>", apiHandler1.GetPoll)
	app.Post("/votes", apiHandlerVotes.AddVote)
 {
    "voteid": 1,
    "voterid": 14,
    "pollid": 1,
    "votevalue": 2

}
	app.Get("/votes", apiHandlerVotes.ListAllVotes)
	app.Get("/votes/:id<uint>", apiHandlerVotes.GetVote)


Payload for post:

