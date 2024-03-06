SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the voter executable"
	@echo "	   run					Run the voter program from code"
	@echo "	   run-bin				Run the voter executable"
	@echo "	   load-db				Add sample data via curl"
	@echo "	   get-by-id			Get a voter by id pass id=<id> on command line"
	@echo "	   get-all				Get all voters"
	@echo "	   update-15			Update record 15, pass a new lastname in using lastname=<lastname> on command line"
	@echo "	   delete-all			Delete all voters"
	@echo "	   delete-by-id			Delete a voter by id pass id=<id> on command line"
	@echo "	   build-amd64-linux	Build amd64/Linux executable"
	@echo "	   build-arm64-linux	Build arm64/Linux executable"



.PHONY: build
build:
	go build .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./todo-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./todo-linux-arm64 .

	
.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./todo

.PHONY: restore-db
restore-db:
	(cp ./data/todo.json.bak ./data/todo.json)

.PHONY: restore-db-windows
restore-db-windows:
	(copy.\data\todo.json.bak .\data\todo.json)

.PHONY: load-db
load-db:
	curl -d '{ "voterid": 14, "firstname": "William", "lastname": "Torning", "VoteHistory": [{"pollid": 1,"voteid": 1,"VoteDate":"2023-07-25T19:10:34.811997-04:00"},{"pollid": 2,"voteid": 2,"VoteDate":"2023-08-25T19:10:34.811997-04:00"}]} ' -H "Content-Type: application/json" -X POST http://localhost:1081/voter 
	curl -d '{ "voterid": 15, "firstname": "Test", "lastname": "Test2", "VoteHistory": [{"pollid": 2,"voteid": 3,"VoteDate":"2022-07-25T19:10:34.811997-04:00"},{"pollid": 3,"voteid": 4,"VoteDate":"2022-08-25T19:10:34.811997-04:00"}]} ' -H "Content-Type: application/json" -X POST http://localhost:1081/voter 

.PHONY: update-2
update-2:
	curl -d '{ "voterid": 15, "firstname": "Test", "lastname": "$(lastname)" }' -H "Content-Type: application/json" -X PUT http://localhost:1080/voter 

.PHONY: get-by-id
get-by-id:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voter/$(id) 

.PHONY: get-all
get-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voter

.PHONY: delete-all
delete-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1080/voter 

.PHONY: delete-by-id
delete-by-id:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1080/voter/$(id) 

