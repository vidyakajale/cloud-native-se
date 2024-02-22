## Voter API

This is a describes application showing API's for voter-api


**Folders:**                                         
    api   :    voter-api-handler                  
    voters:              
    poll  :                    
    votes :                  

    

**API's:**

**ListAllVoters:**              
        localhost:1080/voter  (GET)                 
        app.Get("/voter", apiHandler.ListAllVoters)                         
        Test case: Passed              

   
   
**Add Voter:**               
localhost:1080/voter   (POST)               
app.Post("/voter", apiHandler.AddVoter)                       
Test cases :Passed                  
Payload:                     
                 {       
                    "voterid": 14,        
                    "firstname": "William",          
                    "lastname": "Torning",          
                    "VoteHistory": [{"pollid": 1,"voteid": 1,"VoteDate":"2023-07-25T19:10:34.811997-04:00"},{"pollid": 2,"voteid": 2,"VoteDate":"2023-08-25T19:10:34.811997-04:00"}]        
                }                    


                
**Get Voter by id:**                      
localhost:1080/voter/{id}   (GET)           
app.Get("/voter/:id<uint>", apiHandler.GetVoter)        
Test case: Passed              
	
 
 
**Add Voter by id:**                      
localhost:1080/voter{id}/poll/{pollid}  (POST)                      
app.Post("/voter/:id<uint>/poll/:pollid<uint>", apiHandler.AddVoterPoll)                      
Paylpad:                      
                {                      
                        "pollid": 2,                      
                        "voteid": 45                      
                }                      


**Get voter by id by pollid:**                         
localhost:1080/voter{id}/poll/{pollid}  (GET)                      
app.Get("/voter/:id<uint>/poll/:pollid<uint>", apiHandler.GetVoterPoll)                      
	

      
**Get voter health:**                       
localhost:1080/voters/health  (GET)                      
app.Get("/voters/health", apiHandler.HealthCheck)                      
	


**Add Poll:**                      
app.Post("/poll", apiHandler1.AddPoll)                      
localhost:1080/poll  (POST)                      
Payload:                      
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

 
	
**Get polls:**                      
 app.Get("/poll", apiHandler1.ListAllPolls)                      
 localhost:1080/poll  (GET)                      

                      
                      
**Get poll by id:**                      
app.Get("/poll/:id<uint>", apiHandler1.GetPoll)                      
localhost:1080/poll/{id}  (GET)                      



**Add Votes:**                       
app.Post("/votes", apiHandlerVotes.AddVote)                      
localhost:1080/votes  (POST)                      
Payload:                      
		 `{                      
		    "voteid": 1,                      
		    "voterid": 14,                      
		    "pollid": 1,                      
		    "votevalue": 2                      
		}                      
`


**Get Votes:**                      
app.Get("/votes", apiHandlerVotes.ListAllVotes)                      
localhost:1080/votes  (GET)                      

 
**Get votes by id:**                      
 app.Get("/votes/:id<uint>", apiHandlerVotes.GetVote)                      
 localhost:1080/votes/{id}  (GET)                      

                      
                      
**Test cases Run:**                      

<img width="1230" alt="Screenshot 2024-02-21 at 9 41 14 PM" src="https://github.com/vidyakajale/cloud-native-se/assets/74523779/cb3a20ff-d489-4c91-89fa-05624dd50a1d">






