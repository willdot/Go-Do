# Go-Do
Go-Do is Daily Do but written with Go Microservices


## Usage

Run docker-compose up

When all services are up (the user and task service may take longer as they wait for the Cassandra service to be running first), launch Postman and start making requests.

Url : http://localhost:8080/rpc
Method: Post

### Auth service

The auth service allows users to create, login and change passwords. 

#### Create
Body: 
``` json
{
	"service" : "go_do.auth",
	"method" : "Auth.Create",
	"request" : {
		"name" : "will",
		"company" : "Go Do",
		"email" : "will@email.com",
		"password" : "password"
	}
}
```

This returns the same data back, apart from the password is hashed.

#### Login
Body:
```json
{
	"service" : "go_do.auth",
	"method" : "Auth.Auth",
	"request" : {
		"email" : "will@email.com",
		"password" : "password"
	}
}
```
This returns a JWT that needs to be used for any other requests.

#### Change password
Body:
```json
{
	"service" : "go_do.auth",
	"method" : "Auth.ChangePassword",
	"request" : {
		"email" : "will@email.com",
		"oldPassword" : "password",
		"newPassword" : "password1"
	}
}
```
This returns a new JWT to use.

### Task service

The task service allows users to Create, Get, Complete, Update or change Daily Do status.

#### Create
Header:
    Token: {JWT from Auth service}
Body:
```json
{
	"service" : "go_do.task",
	"method" : "TaskService.Create",
	"request" : {
		"title" : "First test",
		"description" : "fingers crossed",
		"dailyDo" : true
	}
}
```

This returns a Task with some added info such as created date.

#### Get
Header:
    Token: {JWT from Auth service}
Body:
```json
{
	"service" : "go_do.task",
	"method" : "TaskService.Get",
	"request" : {
	}
}
```

This returns an array of task data.

// TODO: Update, Complete and Change Daily Do Status
