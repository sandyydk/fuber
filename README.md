# fuber
A cab service prototype in Golang

A backend RESTful API to simulate cab booking. Has three endpoints which do them. 
1. /getcabs - lists all avaialbe cabs (GET)
2. /beginride - returns the nearest available cab to the user. (POST)
3. /endride - ends a currently active ride (POST)

Response models are being stored in models.go. Handlers contains the code whioch handles the requests. Each major structure like that of user, 
location and cab have their own packages which can be expanded later to add more precision. Includes comments which explain how it is done. Comes with test cases written in Golang 
which have _test suffixes in files names like main_test.go. Can test them by 'cd' ing to the folder and run 'go test'.

The API service can be run by 'go run main.go' in the root folder. It will run in default 3002 port of the localhost. Can use POSTMAN or fiddler to run the tests. 

For more info, contact sandyethadka@gmail.com
