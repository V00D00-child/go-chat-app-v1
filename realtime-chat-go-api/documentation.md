# Documentation Chattie REST API
06/12/19 *  Contributors Idris Bowman

## Requirements
URL: https://api.chattie.com/v1/
All request made to the API require an auth token.
All request to the API with be logged to a JSON file with a timestamp
All user will be save to Mysql database
All messages will be saved to elasticsearch
All error message will be saved to elasticsearch

## Technology versions 
Golang -
mariadb -
elasticsearch -
Angular -  

## Resources
GET/auth - Retrieve auth token (201,401)
GET/users - Retrieve all users (200,401)
POST/users -Create a new user (201,401)
DELETE/user -Remove all users (204,401)

GET/users/1 -	Retrieve the details for user 1(201,401)
PUT/users/1 - Update the details of user 1 if it exists (200,401)
DELETE/users/1 - Remove user 1 (204,401)
POST/users/1 -Method not allowed (405)
PUT/users/1/permissions - Update the details of user 1 credentials (200,401)

GET/messages -Retrieve the all the messages (201,401)
POST/users/1/messages - Create a new message for user 1 (201,401)
GET/users/1/message - Retrieve all messages for user 1(201,401)
PUT/users/1/message - Bulk update of messages for user 1 if it exists (200,401)
DELETE/users/1/message - Remove all messages for user 1 (204,401)
GET/messages/errors - Retrieve the all the error messages (201,401)


## Example request:
GET https://api.chattie.com/v1/user/12

## HTTP status Codes
200 – OK – Everything is working
201 – OK – New resource has been created
204 – OK – The resource was successfully deleted
400 – Bad Request – The request was invalid or cannot be served. The exact error should be explained 
in the error payload. E.g. „The JSON is not valid“.
401 – Unauthorized – The request requires an user authentication
403 – Forbidden – The server understood the request, but is refusing it or the access is not allowed.
404 – Not found – There is no resource behind the URI.
405 - Method Not Allowed
500- Internal Server Error is a general http status code that means something has gone wrong on the website's server

## API Server Error Payload
```json
{
    "message":"Describe the error",
    "type":"What type of error is it",
    "code":"HTTP status code",
    "trace_id": "each error will have a ID(use this for metrics)"
}
```

## Chat Message Payload
```json
{
    "body":"Hey did you catch the game last night",
    "type":"message from chat app",
    "user_id":"23",
    "receivedTimeEpoch":23232432
}
```

## API  access logging capture
```json
{
  "Date": "string",
  "DateEpoch": 600004,
  "Method": "string",
  "URI": "string",
  "SrcIP": "net.IP",
  "SrcPort": "string",
  "Host": "string",
  "Status": 23,
  "UserAgent": "string",
  "RequestHeader": "http.Header"
}
```

## Paging -Coming in v2
GET https://api.chattie.com/v2/users?offset=10&limit=5

