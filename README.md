Execute following commands to run the in local
// Navigate to the folder containing the code

Step-1: go build -o api hello.go

Step-2: sam local start-api

Step-3: curl http://localhost:3000/ping curl http://localhost:3000/health
