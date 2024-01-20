# Task reminder application

This web application enables users to create a task reminder for a specific contact that will remind them of the task. This application can be run by using the command, `go run taskreminder.go --configfile "config.json"`
Root API for this application is "http://localhost:port/api/v1/task/", where port can be configured in the config.json file
There are currently five APIs implemented,
### /health :: which checks if the server is connected to the database.
### /create (POST) :: this creates a task 
### /getAllTasks (GET) :: this fetches all the task including the ones that are pending, completed and expired of a user.
### /getPendingTasks (GET) :: this fetches all the pending tasks of a user
### /updateTask (PUT) :: this updates a particular task details including status using task id
### /deleteTask/id={id} :: this deletes a task based on the given id 
 
An example of I/P curl command is,
curl --location --request GET 'http://localhost:1104/api/v1/task/getAllTasks' \
--header 'Content-Type: application/json' \
--data '{
    "createdby": "user"
}'

The web server is written in golang using chi framework and postgres database. This codebase follows the clean architecture. Clean architecture is a software design philosophy that emphasizes separation of concerns, maintainability, and testability by organizing code into distinct layers, each with its specific responsibility. The architecture encourages dependency inversion, making the inner layers independent of the outer layers.
The main file is taskreminder.go.
The codebase is splitted into multiple subfolders. Below explains what each subfolder contains,
### config - contains code to parse the config.json file.
### database - contains subfolders for each database. Here I have implemented for postgres. This contains files that create database instance, gets the created instance (if not there creates new instance) and clears the instance.
### handler - contains the routing handler functions. Each handler has multiple subfolder. TaskHandler subfolder has multiple files that implements task api routes and blueprint of routes present.
### logs - contains the log files
### models - contains request and response models for the APIs.
### schema - contains the schema of the database for the task table.
### server - containes the files related to setting up, configuring and starting the server.
### services - contains the business logic of the task manager application. It contains two subfolders,
    1. tasksvc - contains files with implementation of task apis and db queries.
    2. taskchecksvc - contains files that start a goroutine to move all the pending task to expired state after their duedate has passed
### utils - contains util functions for struct validation and log configuration.

### postman-screenshots - contains the api testing screenshots.

