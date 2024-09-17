// models/user.go
package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*

Add user:

curl -X POST -H "Content-Type: application/json" \
-d '{"name":"John Doe","email":"john@example.com"}' \
http://localhost:8000/api/users



Parameters:
-X POST: Specifies the HTTP method as POST.
-H "Content-Type: application/json": Sets the content type to JSON.
-d '{...}': The JSON data to send in the request body.



updata user :
curl -X PUT -H "Content-Type: application/json" \
-d '{"name":"John Smith","email":"johnsmith@example.com"}' \
http://localhost:8000/api/users/1


delet user:

curl -X DELETE http://localhost:8000/api/users/1

extra....:
Interpreting Responses and Status Codes
200 OK: The request was successful.
201 Created: A new resource was created successfully.
204 No Content: The request was successful; no content to return.
400 Bad Request: The request could not be understood by the server due to malformed syntax.
404 Not Found: The requested resource does not exist.
500 Internal Server Error: The server encountered an unexpected condition.




x
*/
